package main

import (
	"io"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"

	"encoding/xml"
	"fmt"

	"os"

	"main.go/xml_types"
)

const c_cue_zero = "00:00:00"

type FileDialog struct {
}

type MarkerCue struct {
	TimeCue string
	Text    string
}

func duration_tostring(t_duration time.Duration) string {
	//t_duration := time.Duration(t_seconds * 1e9)
	return fmt.Sprintf("%02d:%02d:%02d", int(t_duration.Hours()), int(t_duration.Minutes())%60, int(t_duration.Seconds())%60)
}

func decode_ticks(ticks string) (string, time.Duration, error) {
	if ticks == "0" || ticks == "0s" {
		return c_cue_zero, time.Duration(0), nil
	}
	//sample ticks = "53873820/24000s"
	Splits := strings.Split(ticks, "/")
	if len(Splits) > 1 {
		ticks_a, _ := strconv.Atoi(Splits[0])
		ticks_b, _ := strconv.Atoi(strings.TrimSuffix(Splits[1], "s"))
		t_seconds := float64(ticks_a / ticks_b)
		t_duration := time.Duration(t_seconds * 1e9)
		s_time := duration_tostring(t_duration)
		//s_time := fmt.Sprintf("%02d:%02d:%02d", int(t_duration.Hours()), int(t_duration.Minutes())%60, int(t_duration.Seconds())%60)

		/*
			fmt.Println("Parsed ticks: ", ticks_a, "/", ticks_b)
			fmt.Println("Parsed seconds: ", t_seconds)
			fmt.Println("Parsed duration: ", t_duration)
			fmt.Println("Parsed time: ", s_time)
		*/

		return s_time, t_duration, nil
	} else {
		ticks_b, _ := strconv.Atoi(strings.TrimSuffix(ticks, "s"))
		t_seconds := float64(ticks_b)
		t_duration := time.Duration(t_seconds * 1e9)
		s_time := duration_tostring(t_duration)
		return s_time, t_duration, nil
	}

	//return c_cue_zero, time.Duration(0), errors.New("Cannot convert, expected format: '53873820/24000s' ")
}

func convert_ticks(ticks string) string {
	decoded_ticks, _, err := decode_ticks(ticks)
	if err == nil {
		return decoded_ticks

	} else {
		return c_cue_zero
	}

}

func printSlice(s []MarkerCue) (text_out string) {
	if len(s) == 0 {
		return ""
	}
	//fmt.Println(s[0].TimeCue)
	return (s[0].TimeCue + " " + s[0].Text + "\n" + printSlice(s[1:]))

}

func NewFileOpen(callback func(fyne.URIReadCloser, error), parent fyne.Window) *FileDialog {
	return nil
}

func OpenXML(file_path string) []byte {
	// Open our xmlFile

	//absPath, _ := filepath.Abs("./xml_types/sample_fcpx.xml")

	absPath := ""

	if file_path != "" {
		absPath = file_path
	} else {
		filename, err := dialog.File().Filter("FCPX XML file", "fcpxmld", "fcpxml", "xml").Title("Open FCPX xml").Load()
		if filename != "" && err == nil {
			absPath = filename
		}
	}

	fmt.Println(absPath)
	fmt.Println(os.Getwd())

	xmlFile, err := os.Open(absPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		// read our opened xmlFile as a byte array.
		byteValue, _ := io.ReadAll(xmlFile)
		//fmt.Println(xmlFile)
		return byteValue
	}

	//fmt.Println("Successfully Opened xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	return nil

}

func Decode_Markers(xml_byte []byte) string {

	if xml_byte == nil {
		return ""
	}

	var fcpxResponse2 = &xml_types.Fcpxml_Library{}
	if err := xml.Unmarshal([]byte(xml_byte), fcpxResponse2); err != nil {
		panic(err)
	} else {
		// fmt.Println(fcpxResponse2)
	}

	//var spine1 = fcpxResponse1.Library.Event.Project.Sequence.Spine
	var spine2 = fcpxResponse2.Library.Event.Project.Sequence.Spine

	//for i := 0; i < len(spine1.AssetClip); i++ {
	//	fmt.Println("User Type: " + spine1.AssetClip[i].Offset)
	//}

	var markers1 []MarkerCue

	for i := 0; i < len(spine2.AssetClip); i++ {
		var assetClipNext = spine2.AssetClip[i]
		cOffset_str, cOffset_dur, err := decode_ticks(assetClipNext.Offset)
		if err != nil {
			fmt.Println(i, "Clip Offset error:	", assetClipNext.Offset)
		}

		var clip_start_dur time.Duration = 0
		clip_start_str := c_cue_zero

		if len(assetClipNext.ChapterMarker) > 0 {
			fmt.Println()
			fmt.Println(i, "Clip Name:	", assetClipNext.Name)
			fmt.Println(i, "Clip Offset:	", cOffset_str)
			fmt.Println(i, "Clip Duration:	", convert_ticks(assetClipNext.Keyword.Duration))

			clip_start_str, clip_start_dur, err = decode_ticks(assetClipNext.Start)
			if err == nil && clip_start_dur != 0 {
				fmt.Println(i, "Clip Start	", clip_start_str)
			}
		}

		for j := 0; j < len(assetClipNext.ChapterMarker); j++ {
			fmt.Println(i, j, "Marker text:	", assetClipNext.ChapterMarker[j].Value)

			marker_str, marker_dur, err := decode_ticks(assetClipNext.ChapterMarker[j].Start)
			if err == nil {

				marker_dur += cOffset_dur - clip_start_dur

				fmt.Println(i, j, "Marker Start:	", marker_str)
				fmt.Println(i, j, "Calculated Duration:	", duration_tostring(marker_dur))

				markers1 = append(markers1, MarkerCue{TimeCue: duration_tostring(marker_dur), Text: assetClipNext.ChapterMarker[j].Value})

			}

		}
	}

	if markers1[0].TimeCue != c_cue_zero {
		//marker_first := []MarkerCue{{TimeCue: c_cue_zero, Text: "Start"}}
		markers1 = append([]MarkerCue{{TimeCue: c_cue_zero, Text: "Start"}}, markers1...)
		fmt.Println(markers1)
		//markers1 = append(, markers1)
	}

	return printSlice(markers1)

}

func ShowApp(disp_text string) {
	//fmt.Println(markers1)
	fmt.Println(disp_text)

	myApp := app.New()
	myWindow := myApp.NewWindow("FCPX Markers")

	txtBound := binding.NewString()
	txtWid := widget.NewEntryWithData(txtBound)
	txtWid.MultiLine = true

	// we can disable the Entry field so the user can't modify the text:
	txtWid.Disabled()

	txtBound.Set(disp_text)

	/*
		go func() {
			for {
				txtBound.Set(time.Now().Format("2006-01-02\n15:04:05\nMST -0700"))
				time.Sleep(time.Second)
			}
		}()
	*/

	content := container.NewBorder(nil, nil, nil, nil, txtWid)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.Size{500, 300})

	//absPath, _ := filepath.Abs("./xml_types/sample_fcpx.xml")

	//fmt.Println(absPath)
	//fmt.Println(os.Getwd())

	//ok := dialog.Message("%s", "Do you want to continue?").Title("Are you sure?").YesNo()

	//if ok == false {

	//os.Exit(0)

	//}

	myWindow.ShowAndRun()
}

func main() {

	// Open our xmlFile
	xmlfile_asbyte := OpenXML("")

	text_out := Decode_Markers(xmlfile_asbyte)

	ShowApp(text_out)

	os.Exit(0)
}
