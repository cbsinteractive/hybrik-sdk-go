package hybrik

import "time"

type JobResultResponse struct {
	Errors []JobResultError `json:"errors"`
	Job    JobResultSummary `json:"job"`
	Tasks  []Tasks          `json:"tasks"`
}

type JobResultSummary struct {
	ID                 int       `json:"id"`
	JobClass           int       `json:"job_class"`
	IsAPIJob           int       `json:"is_api_job"`
	Priority           int       `json:"priority"`
	CreationTime       time.Time `json:"creation_time"`
	ExpirationTime     time.Time `json:"expiration_time"`
	HintActiveTimeSec  int       `json:"hint_active_time_sec"`
	HintMachineTimeSec int       `json:"hint_machine_time_sec"`
	LastTimesUpdated   time.Time `json:"last_times_updated"`
	SubscriptionKey    string    `json:"subscription_key"`
	Flags              int       `json:"flags"`
	Status             string    `json:"status"`
	RenderStatus       string    `json:"render_status"`
	TaskCount          int       `json:"task_count"`
	Progress           int       `json:"progress"`
	Name               string    `json:"name"`
	FirstStarted       time.Time `json:"first_started"`
	LastCompleted      time.Time `json:"last_completed"`
}

type JobResultError struct {
	TaskInstanceID   int         `json:"task_instance_id"`
	RetryNr          int         `json:"retry_nr"`
	MachineFetcherID interface{} `json:"machine_fetcher_id"`
	ResultDefine     string      `json:"result_define"`
	Message          string      `json:"message"`
	Details          interface{} `json:"details"`
	Diagnostic       interface{} `json:"diagnostic"`
	RecoverableError int         `json:"recoverable_error"`
	Assigned         interface{} `json:"assigned"`
	ResultCommitted  time.Time   `json:"result_committed"`
}

type ExecutionResultDetails struct {
	MachineID int `json:"machine_id"`
	ServiceID int `json:"service_id"`
	TaskID    int `json:"task_id"`
}

type LocationResult struct {
	Path            string `json:"path"`
	StorageProvider string `json:"storage_provider"`
}

type AssetComponentsResult struct {
	Kind         string                     `json:"kind"`
	Name         string                     `json:"name"`
	Options      TranscodeTaskOptionsResult `json:"options"`
	ComponentUID string                     `json:"component_uid"`
}

type AssetVersionsResult struct {
	Location        LocationResult          `json:"location"`
	AssetComponents []AssetComponentsResult `json:"asset_components"`
	VersionUID      string                  `json:"version_uid"`
}
type AssetResultPayload struct {
	AssetVersions []AssetVersionsResult `json:"asset_versions"`
	Kind          string                `json:"kind"`
}

type ResultPayload struct {
	Kind    string             `json:"kind"`
	Payload AssetResultPayload `json:"payload"`
}

type DocumentResult struct {
	ResultPayload ResultPayload `json:"result_payload"`
	Connector     string        `json:"connector"`
}

type Params struct {
	Location    LocationResult `json:"location"`
	FilePattern string         `json:"file_pattern"`
}
type MezzQCResultPayload struct {
	Module string `json:"module"`
	Params Params `json:"params"`
}

type ResultConfig struct {
	UID     string             `json:"uid"`
	Kind    string             `json:"kind"`
	Payload AssetResultPayload `json:"payload"`
	Name    string             `json:"name"`
}

type PreprocessingResult struct {
	Task TaskTags `json:"task"`
}

type TranscodeTaskResult struct {
	Name string `json:"name"`
}

type TranscodeTaskContainerResult struct {
	Kind string `json:"kind"`
}

type TranscodeTaskVideoResult struct {
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	BitrateMode    string `json:"bitrate_mode"`
	MinBitrateKb   int    `json:"min_bitrate_kb"`
	BitrateKb      int    `json:"bitrate_kb"`
	MaxBitrateKb   int    `json:"max_bitrate_kb"`
	Preset         string `json:"preset"`
	Codec          string `json:"codec"`
	Profile        string `json:"profile"`
	MinGopFrames   int    `json:"min_gop_frames"`
	MaxGopFrames   int    `json:"max_gop_frames"`
	ExactGopFrames int    `json:"exact_gop_frames"`
	InterlaceMode  string `json:"interlace_mode"`
	X265Options    string `json:"x265_options"`
	Vtag           string `json:"vtag"`
	FfmpegArgs     string `json:"ffmpeg_args"`
}

type TranscodeTaskTargetResult struct {
	FilePattern   string                       `json:"file_pattern"`
	Container     TranscodeTaskContainerResult `json:"container"`
	Video         TranscodeTaskVideoResult     `json:"video"`
	ExistingFiles string                       `json:"existing_files"`
}

type SegmentedRenderingResult struct {
	DurationSec int `json:"duration_sec"`
}

type SourcePipeline struct {
	SegmentedRendering SegmentedRenderingResult `json:"segmented_rendering"`
}

type OptionsPipelineResult struct {
	EncoderVersion string `json:"encoder_version"`
}

type TranscodeTaskOptionsResult struct {
	Pipeline       OptionsPipelineResult `json:"pipeline"`
	SourceReadMode string                `json:"source_read_mode"`
}

type TranscodeTaskResultPayload struct {
	Location       LocationResult              `json:"location"`
	Targets        []TranscodeTaskTargetResult `json:"targets"`
	SourcePipeline SourcePipeline              `json:"source_pipeline"`
	Options        TranscodeTaskOptionsResult  `json:"options"`
}

type TranscodeElementResult struct {
	UID     string              `json:"uid"`
	Kind    string              `json:"kind"`
	Task    TranscodeTaskResult `json:"task"`
	Payload AssetResultPayload  `json:"payload"`
}

type Mp4MuxResult struct {
	Enabled     bool   `json:"enabled"`
	FilePattern string `json:"file_pattern"`
	ToolVersion string `json:"tool_version"`
}

type PostTranscodeResult struct {
	Mp4Mux Mp4MuxResult `json:"mp4_mux"`
}

type DoViResultPayload struct {
	Module        string                   `json:"module"`
	Profile       int                      `json:"profile"`
	Location      LocationResult           `json:"location"`
	Preprocessing PreprocessingResult      `json:"preprocessing"`
	Transcodes    []TranscodeElementResult `json:"transcodes"`
	PostTranscode PostTranscodeResult      `json:"post_transcode"`
}

type Tasks struct {
	ID               int              `json:"id"`
	Priority         int              `json:"priority"`
	RetryNr          int              `json:"retry_nr"`
	RetryNrAog       int              `json:"retry_nr_aog"`
	CreationTime     time.Time        `json:"creation_time"`
	MaxRetryCountAog int              `json:"max_retry_count_aog"`
	RelatedAssetID   interface{}      `json:"related_asset_id"`
	Kind             string           `json:"kind"`
	Name             string           `json:"name"`
	RetryCount       int              `json:"retry_count"`
	UID              string           `json:"uid"`
	ElementName      string           `json:"element_name"`
	Status           string           `json:"status"`
	Assigned         time.Time        `json:"assigned"`
	Completed        time.Time        `json:"completed"`
	Documents        []DocumentResult `json:"documents"`
	FetcherID        int              `json:"fetcher_id,omitempty"`
}
