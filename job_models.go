package hybrik

const (
	DoViVesQCVersionDefault    = "0.9.0.9"
	DoViMezzQCVersionDefault   = "2.6.2"
	DoViMP4QCVersionDefault    = "1.1.4"
	DoViMP4MuxerVersionDefault = "1.2.8"
	DoViSDKVersionDefault      = "4.2.1_ga"
)

// CreateJob .
type CreateJob struct {
	Name              string           `json:"name"`
	Payload           CreateJobPayload `json:"payload"`
	Schema            string           `json:"schema,omitempty"`
	Expiration        int              `json:"expiration,omitempty"`
	Priority          int              `json:"priority,omitempty"`
	TaskRetryCount    int              `json:"task_retry:count,omitempty"`
	TaskRetryDelaySec int              `json:"task_retry:delay_sec,omitempty"`
	TaskTags          []string         `json:"task_tags,omitempty"`
	UserTag           string           `json:"user_tag,omitempty"`
}

// CreateJobPayload .
type CreateJobPayload struct {
	Elements    []Element    `json:"elements,omitempty"`
	Connections []Connection `json:"connections,omitempty"`
}

// Element .
type Element struct {
	UID     string              `json:"uid"`
	Kind    string              `json:"kind"`
	Task    *ElementTaskOptions `json:"task,omitempty"`
	Preset  *TranscodePreset    `json:"preset,omitempty"`
	Payload interface{}         `json:"payload"` // Can be of type ElementPayload or LocationTargetPayload
}

// ElementTaskOptions .
type ElementTaskOptions struct {
	Name string `json:"name"`
}

// ElementPayload .
type ElementPayload struct {
	Kind    string       `json:"kind,omitempty"`
	Payload AssetPayload `json:"payload"`
}

// ManifestCreatorPayload .
type ManifestCreatorPayload struct {
	Location    TranscodeLocation `json:"location"`
	FilePattern string            `json:"file_pattern"`
	Kind        string            `json:"kind"`
	UID         string            `json:"uid,omitempty"`
}

// LocationTargetPayload .
type LocationTargetPayload struct {
	Location TranscodeLocation `json:"location"`
	Targets  interface{}       `json:"targets"`
}

// TranscodePayload holds configurations for a transcode task
type TranscodePayload struct {
	LocationTargetPayload
	SourcePipeline TranscodeSourcePipeline `json:"source_pipeline,omitempty"`
	Options        *TranscodeTaskOptions   `json:"options,omitempty"`
}

type TaskTags struct {
	Tags []string `json:"tags,omitempty"`
}

type DolbyVisionTaskPayload struct {
	Module        string            `json:"module"`
	Profile       int               `json:"profile"`
	MezzanineQC   DoViMezzanineQC   `json:"mezzanine_qc,omitempty"`
	NBCPreproc    DoViNBCPreproc    `json:"nbc_preproc,omitempty"`
	Transcodes    []Element         `json:"transcodes"`
	PostTranscode DoViPostTranscode `json:"post_transcode,omitempty"`
}

// DoViMezzanineQC holds mezz qc config options
type DoViMezzanineQC struct {
	Enabled     bool              `json:"enabled"`
	Location    TranscodeLocation `json:"location"`
	Task        TaskTags          `json:"task,omitempty"`
	FilePattern string            `json:"file_pattern"`
	ToolVersion string            `json:"tool_version"`
}

// DoViNBCPreproc holds the DolbyVision pre-processor config
type DoViNBCPreproc struct {
	Task           TaskTags                 `json:"task,omitempty"`
	Location       TranscodeLocation        `json:"location"`
	SDKVersion     string                   `json:"dovi_sdk_version"`
	NumTasks       string                   `json:"num_tasks"`
	IntervalLength int                      `json:"interval_length"`
	CLIOptions     DoViNBCPreprocCLIOptions `json:"cli_options,omitempty"`
}

// DoViNBCPreprocCLIOptions contains command line options for the nbc_preproc task for DolbyVision
type DoViNBCPreprocCLIOptions struct {
	InputEDRSize   string `json:"inputEDRSize,omitempty"`
	InputEDRAspect string `json:"inputEDRAspect,omitempty"`
	InputEDRPad    string `json:"inputEDRPad,omitempty"`
	InputEDRCrop   string `json:"inputEDRCrop,omitempty"`
}

// DoViPostTranscode holds configuration for the DolbyVision post-transcode settings
type DoViPostTranscode struct {
	Task             TaskTags             `json:"task,omitempty"`
	VESMux           DoViVESMux           `json:"ves_mux,omitempty"`
	MetadataPostProc DoViMetadataPostProc `json:"metadata_postproc,omitempty"`
	MP4Mux           DoViMP4Mux           `json:"mp4_mux,omitempty"`
}

// DoViVESMux configures settings for the VES muxing post-transcode
type DoViVESMux struct {
	Enabled     bool              `json:"enabled"`
	Location    TranscodeLocation `json:"location"`
	FilePattern string            `json:"file_pattern"`
	SDKVersion  string            `json:"dovi_sdk_version"`
}

// DoViMetadataPostProc configures settings for the metadata post processing after a DolbyVision transcode
type DoViMetadataPostProc struct {
	Enabled     bool              `json:"enabled"`
	Location    TranscodeLocation `json:"location"`
	FilePattern string            `json:"file_pattern"`
	SDKVersion  string            `json:"dovi_sdk_version"`
	QCSettings  DoViQCSettings    `json:"qc,omitempty"`
}

// DoViQCSettings holds settings for the post transcode DoVi metadata qc job
type DoViQCSettings struct {
	Enabled     bool              `json:"enabled"`
	ToolVersion string            `json:"tool_version"`
	Location    TranscodeLocation `json:"location"`
	FilePattern string            `json:"file_pattern"`
}

// DoViMP4Mux holds settings for the DolbyVision mp4 muxer
type DoViMP4Mux struct {
	Enabled            bool                         `json:"enabled"`
	Location           TranscodeLocation            `json:"location"`
	FilePattern        string                       `json:"file_pattern"`
	ToolVersion        string                       `json:"tool_version"`
	CopySourceStartPTS bool                         `json:"copy_source_start_pts"`
	QCSettings         DoViQCSettings               `json:"qc,omitempty"`
	ElementaryStreams  []DoViMP4MuxElementaryStream `json:"elementary_streams,omitempty"`
}

// DoViMP4MuxElementaryStream holds settings for streams during a mux operation
type DoViMP4MuxElementaryStream struct {
	AssetURL        AssetURL              `json:"asset_url,omitempty"`
	ExtractAudio    bool                  `json:"extract_audio"`
	ExtractLocation TranscodeLocation     `json:"extract_location,omitempty"`
	ExtractTask     DoViMP4MuxExtractTask `json:"extract_task,omitempty"`
}

// DoViMP4MuxExtractTask hold configurations for extracting data from elementary streams
type DoViMP4MuxExtractTask struct {
	RetryMethod string   `json:"retry_method,omitempty"`
	Retry       Retry    `json:"retry,omitempty"`
	Name        string   `json:"name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// Retry .
type Retry struct {
	Count    int `json:"count,omitempty"`
	DelaySec int `json:"delay_sec,omitempty"`
}

// AssetURL .
type AssetURL struct {
	StorageProvider string `json:"storage_provider,omitempty"`
	URL             string `json:"url,omitempty"`
}

// TranscodeSourcePipeline allows the modification of the source prior to beginning the transcode
type TranscodeSourcePipeline struct {
	// Segmented rendering parameters.
	SegmentedRendering *SegmentedRendering `json:"segmented_rendering,omitempty"`

	// The FFmpeg source string to be applied to the source file. Use {source_url} within this string
	// to insert the source file name(s).
	FfmpegSourceArgs string `json:"ffmpeg_source_args,omitempty"`

	// SourcePipeline options
	Options TranscodeSourcePipelineOpts `json:"options,omitempty"`

	// Use accelerated Apple ProRes decoder.
	EnableAcceleratedProres bool `json:"accelerated_prores,omitempty"`

	// Defines the level of complexity allowed when using a manifest as a source.
	// Valid values are: 'simple', 'reject_complex' or 'reject_master_playlist'
	DecodeStrategy string `json:"manifest_decode_strategy,omitempty"`

	// The dithering algorithm to use for color conversions.
	// Valid values are:
	// 'none', 'bayer', 'ed', 'a_dither' or 'x_dither'
	ChromaDitherAlgorithm string `json:"chroma_dither_algorithm,omitempty"`

	// The type of function to be used in scaling operations.
	Scaler TranscodeSourcePipelineScaler `json:"scaler,omitempty"`
}

// SegmentedRendering holds segmented rendering parameters
type SegmentedRendering struct {
	// Duration (in seconds) of a segment in segment encode mode. Minimum: 1
	Duration int `json:"duration_sec,omitempty"`

	// Duration (in seconds) to look for a dominant previous or following scene change. Note that
	// the segment duration can then be up to duration_sec + scene_changes_search_duration_sec long.
	SceneChangeSearchDuration int `json:"scene_changes_search_duration_sec,omitempty"`

	// Total number of segments
	NumTotalSegments int `json:"total_segments,omitempty"`

	// Combiner will merge and re-stripe transport streams
	EnableStrictCFR bool `json:"strict_cfr,omitempty"`

	// Timebase offset to be used by the muxer
	MuxTimebaseOffset int `json:"mux_offset_otb,omitempty"`
}

// TranscodeSourcePipelineOpts are extra options you can add to a transcode source pipeline
type TranscodeSourcePipelineOpts struct {
	// Forces Fixed Frame Rate - even if the source file is detected as a variable frame rate source,
	// treat it as a fixed framerate source.
	ForceFixedFrameRate bool `json:"force_ffr,omitempty"`

	// Sets the maximum time for waiting to access the source data. This can be used to handle data that is in transit.
	SourceFetchTimeout int `json:"wait_for_source_timeout_sec,omitempty"`

	// The maximum number of decode errors to allow. Normally, decode errors cause job failure, but
	// there can be situations where a more flexible approach is desired.
	MaxDecodeErrors int `json:"max_decode_errors,omitempty"`

	// The maximum number of sequential errors to allow during decode. This can be used in combination with
	// max_decode_errors to set bounds on allowable errors in the source.
	MaxSequentialDecodeErrors int `json:"max_sequential_decode_errors,omitempty"`

	// Certain files may generate A/V sync issues when rewinding, for example after a pre-analysis. This will enforce
	// a reset instead of rewinding.
	DisableRewind bool `json:"no_rewind,omitempty"`

	// Certain files should never be seeked because of potentially occurring precision issues.
	DisableSeek bool `json:"no_seek,omitempty"`

	// Allows files to be loaded in low latency mode, meaning that there will be no analysis at startup.
	DisableAnalysis bool `json:"low_latency,omitempty"`

	// If a render node is allowed to cache this file, this will set the Time To Live (ttl). If not set
	// (or set to 0) the file will not be cached but re-obtained whenever required.
	SourceCacheTTL int `json:"cache_ttl,omitempty"`

	// If this is set to true, the file is considered a manifest. The media files referred to in the
	// manifest will be taken as the real source.
	ResolveManifest bool `json:"resolve_manifest,omitempty"`

	// If this is set to true, the source is considered starting with PTS 0 regardless of the actual PTS.
	ResetPTS bool `json:"auto_offset_sources,omitempty"`
}

// TranscodeSourcePipelineScaler holds scaling parameters to be applied before transcoding
type TranscodeSourcePipelineScaler struct {
	// The type of scaling to be applied.
	// Valid values: 'default' or 'zscale'
	Kind string `json:"kind,omitempty"`

	// The configuration string to be used with the specified scaling function.
	Config string `json:"config_string,omitempty"`

	// Always use the specified scaling function.
	ApplyAlways bool `json:"apply_always,omitempty"`
}

// TranscodePreset .
type TranscodePreset struct {
	Key string `json:"key"`
}

// TranscodeLocationTarget .
type TranscodeLocationTarget struct {
	FilePattern   string                   `json:"file_pattern"`
	ExistingFiles string                   `json:"existing_files,omitempty"`
	Container     TranscodeTargetContainer `json:"container,omitempty"`
	Location      *TranscodeLocation       `json:"location,omitempty"`
}

// TranscodeTargetContainer .
type TranscodeTargetContainer struct {
	Kind            string `json:"kind"`
	SegmentDuration uint   `json:"segment_duration,omitempty"`
}

// AssetPayload .
type AssetPayload struct {
	StorageProvider string `json:"storage_provider,omitempty"`

	URL string `json:"url,omitempty"`
}

// TranscodeLocation .
type TranscodeLocation struct {
	StorageProvider string `json:"storage_provider,omitempty"`
	Path            string `json:"path,omitempty"`
}

//TranscodeTarget .
type TranscodeTarget struct {
	FilePattern   string                   `json:"file_pattern"`
	ExistingFiles string                   `json:"existing_files"`
	Container     TranscodeContainer       `json:"container"`
	Video         map[string]interface{}   `json:"video"`
	Audio         []map[string]interface{} `json:"audio"`
}

// TranscodeContainer .
type TranscodeContainer struct {
	Kind              string `json:"kind"`
	HEVCCodecIDPrefix string `json:"hevc_codec_id_prefix,omitempty"`
}

// Connection .
type Connection struct {
	From []ConnectionFrom `json:"from,omitempty"`
	To   ConnectionTo     `json:"to,omitempty"`
}

// ConnectionFrom .
type ConnectionFrom struct {
	Element string `json:"element,omitempty"`
}

// ConnectionTo .
type ConnectionTo struct {
	Success []ToSuccess `json:"success,omitempty"`
	Error   []ToError   `json:"error,omitempty"`
}

// ToSuccess .
type ToSuccess struct {
	Element string `json:"element,omitempty"`
}

// ToError .
type ToError struct {
	Element string `json:"element,omitempty"`
}
