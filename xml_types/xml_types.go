package xml_types

// build by https://xml-to-go.github.io

import (
	"encoding/xml"
)

type Fcpxml_Full struct {
	XMLName   xml.Name `xml:"fcpxml"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"version,attr"`
	Resources struct {
		Text   string `xml:",chardata"`
		Format []struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id,attr"`
			Name          string `xml:"name,attr"`
			FrameDuration string `xml:"frameDuration,attr"`
			Width         string `xml:"width,attr"`
			Height        string `xml:"height,attr"`
			ColorSpace    string `xml:"colorSpace,attr"`
		} `xml:"format"`
		Asset []struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id,attr"`
			Name          string `xml:"name,attr"`
			Uid           string `xml:"uid,attr"`
			Start         string `xml:"start,attr"`
			Duration      string `xml:"duration,attr"`
			HasVideo      string `xml:"hasVideo,attr"`
			Format        string `xml:"format,attr"`
			HasAudio      string `xml:"hasAudio,attr"`
			VideoSources  string `xml:"videoSources,attr"`
			AudioSources  string `xml:"audioSources,attr"`
			AudioChannels string `xml:"audioChannels,attr"`
			AudioRate     string `xml:"audioRate,attr"`
			MediaRep      struct {
				Text     string `xml:",chardata"`
				Kind     string `xml:"kind,attr"`
				Sig      string `xml:"sig,attr"`
				Src      string `xml:"src,attr"`
				Bookmark string `xml:"bookmark"`
			} `xml:"media-rep"`
			Metadata struct {
				Text string `xml:",chardata"`
				Md   []struct {
					Text  string `xml:",chardata"`
					Key   string `xml:"key,attr"`
					Value string `xml:"value,attr"`
					Array struct {
						Text   string   `xml:",chardata"`
						String []string `xml:"string"`
					} `xml:"array"`
				} `xml:"md"`
			} `xml:"metadata"`
		} `xml:"asset"`
	} `xml:"resources"`
	Library struct {
		Text     string `xml:",chardata"`
		Location string `xml:"location,attr"`
		Event    struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Uid     string `xml:"uid,attr"`
			Project struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Uid      string `xml:"uid,attr"`
				ModDate  string `xml:"modDate,attr"`
				Sequence struct {
					Text        string `xml:",chardata"`
					Format      string `xml:"format,attr"`
					Duration    string `xml:"duration,attr"`
					TcStart     string `xml:"tcStart,attr"`
					TcFormat    string `xml:"tcFormat,attr"`
					AudioLayout string `xml:"audioLayout,attr"`
					AudioRate   string `xml:"audioRate,attr"`
					Spine       struct {
						Text      string `xml:",chardata"`
						AssetClip []struct {
							Text        string `xml:",chardata"`
							Ref         string `xml:"ref,attr"`
							Offset      string `xml:"offset,attr"`
							Name        string `xml:"name,attr"`
							Duration    string `xml:"duration,attr"`
							Format      string `xml:"format,attr"`
							TcFormat    string `xml:"tcFormat,attr"`
							AudioRole   string `xml:"audioRole,attr"`
							ConformRate struct {
								Text         string `xml:",chardata"`
								ScaleEnabled string `xml:"scaleEnabled,attr"`
							} `xml:"conform-rate"`
							Keyword struct {
								Text     string `xml:",chardata"`
								Start    string `xml:"start,attr"`
								Duration string `xml:"duration,attr"`
								Value    string `xml:"value,attr"`
							} `xml:"keyword"`
							AudioChannelSource struct {
								Text  string `xml:",chardata"`
								SrcCh string `xml:"srcCh,attr"`
								Role  string `xml:"role,attr"`
							} `xml:"audio-channel-source"`
							ChapterMarker []struct {
								Text         string `xml:",chardata"`
								Start        string `xml:"start,attr"`
								Duration     string `xml:"duration,attr"`
								Value        string `xml:"value,attr"`
								PosterOffset string `xml:"posterOffset,attr"`
							} `xml:"chapter-marker"`
						} `xml:"asset-clip"`
					} `xml:"spine"`
				} `xml:"sequence"`
			} `xml:"project"`
		} `xml:"event"`
		SmartCollection []struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Match     string `xml:"match,attr"`
			MatchClip struct {
				Text string `xml:",chardata"`
				Rule string `xml:"rule,attr"`
				Type string `xml:"type,attr"`
			} `xml:"match-clip"`
			MatchMedia []struct {
				Text string `xml:",chardata"`
				Rule string `xml:"rule,attr"`
				Type string `xml:"type,attr"`
			} `xml:"match-media"`
			MatchRatings struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"match-ratings"`
		} `xml:"smart-collection"`
	} `xml:"library"`
}

type Fcpxml_Library struct {
	XMLName xml.Name `xml:"fcpxml"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Library struct {
		Text     string `xml:",chardata"`
		Location string `xml:"location,attr"`
		Event    struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Uid     string `xml:"uid,attr"`
			Project struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Uid      string `xml:"uid,attr"`
				ModDate  string `xml:"modDate,attr"`
				Sequence struct {
					Text        string `xml:",chardata"`
					Format      string `xml:"format,attr"`
					Duration    string `xml:"duration,attr"`
					TcStart     string `xml:"tcStart,attr"`
					TcFormat    string `xml:"tcFormat,attr"`
					AudioLayout string `xml:"audioLayout,attr"`
					AudioRate   string `xml:"audioRate,attr"`
					Spine       struct {
						Text      string `xml:",chardata"`
						AssetClip []struct {
							Text        string `xml:",chardata"`
							Ref         string `xml:"ref,attr"`
							Offset      string `xml:"offset,attr"`
							Name        string `xml:"name,attr"`
							Start       string `xml:"start,attr"`
							Duration    string `xml:"duration,attr"`
							Format      string `xml:"format,attr"`
							TcFormat    string `xml:"tcFormat,attr"`
							AudioRole   string `xml:"audioRole,attr"`
							ConformRate struct {
								Text         string `xml:",chardata"`
								ScaleEnabled string `xml:"scaleEnabled,attr"`
							} `xml:"conform-rate"`
							Keyword struct {
								Text     string `xml:",chardata"`
								Start    string `xml:"start,attr"`
								Duration string `xml:"duration,attr"`
								Value    string `xml:"value,attr"`
							} `xml:"keyword"`
							AudioChannelSource struct {
								Text  string `xml:",chardata"`
								SrcCh string `xml:"srcCh,attr"`
								Role  string `xml:"role,attr"`
							} `xml:"audio-channel-source"`
							ChapterMarker []struct {
								Text         string `xml:",chardata"`
								Start        string `xml:"start,attr"`
								Duration     string `xml:"duration,attr"`
								Value        string `xml:"value,attr"`
								PosterOffset string `xml:"posterOffset,attr"`
							} `xml:"chapter-marker"`
						} `xml:"asset-clip"`
					} `xml:"spine"`
				} `xml:"sequence"`
			} `xml:"project"`
		} `xml:"event"`
	} `xml:"library"`
}

type PostResponse struct {
	PostID     string `xml:"Body>postResponse>return>postResult>postId"`
	PostNumber string `xml:"Body>postResponse>return>postNumber"`
}

type ChapterMarker struct {
	Start        string `xml:"start,attr"`
	Duration     string `xml:"duration,attr"`
	Value        string `xml:"value,attr"`
	PosterOffset string `xml:"posterOffset,attr"`
}
type AssetClip struct {
	Ref string `xml:"ref,attr"`
	//ChapterMarkers []ChapterMarker `xml:"chapter-marker"`
}

type Spine struct {
	XMLName    xml.Name    `xml:"spine"`
	AssetClips []AssetClip `xml:"fcpxml>library>event>project>sequence>spine>asset-clip"`
}
