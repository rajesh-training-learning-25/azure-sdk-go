//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// AudioTaskLabel - Defines the possible descriptors for available audio operation responses.
type AudioTaskLabel string

const (
	// AudioTaskLabelTranscribe - Accompanying response data resulted from an audio transcription task.
	AudioTaskLabelTranscribe AudioTaskLabel = "transcribe"
	// AudioTaskLabelTranslate - Accompanying response data resulted from an audio translation task.
	AudioTaskLabelTranslate AudioTaskLabel = "translate"
)

// PossibleAudioTaskLabelValues returns the possible values for the AudioTaskLabel const type.
func PossibleAudioTaskLabelValues() []AudioTaskLabel {
	return []AudioTaskLabel{
		AudioTaskLabelTranscribe,
		AudioTaskLabelTranslate,
	}
}

// AudioTranscriptionFormat - Defines available options for the underlying response format of output transcription information.
type AudioTranscriptionFormat string

const (
	// AudioTranscriptionFormatJSON - Use a response body that is a JSON object containing a single 'text' field for the transcription.
	AudioTranscriptionFormatJSON AudioTranscriptionFormat = "json"
	// AudioTranscriptionFormatSrt - Use a response body that is plain text in SubRip (SRT) format that also includes timing information.
	AudioTranscriptionFormatSrt AudioTranscriptionFormat = "srt"
	// AudioTranscriptionFormatText - Use a response body that is plain text containing the raw, unannotated transcription.
	AudioTranscriptionFormatText AudioTranscriptionFormat = "text"
	// AudioTranscriptionFormatVerboseJSON - Use a response body that is a JSON object containing transcription text along with
	// timing, segments, and other
	// metadata.
	AudioTranscriptionFormatVerboseJSON AudioTranscriptionFormat = "verbose_json"
	// AudioTranscriptionFormatVtt - Use a response body that is plain text in Web Video Text Tracks (VTT) format that also includes
	// timing information.
	AudioTranscriptionFormatVtt AudioTranscriptionFormat = "vtt"
)

// PossibleAudioTranscriptionFormatValues returns the possible values for the AudioTranscriptionFormat const type.
func PossibleAudioTranscriptionFormatValues() []AudioTranscriptionFormat {
	return []AudioTranscriptionFormat{
		AudioTranscriptionFormatJSON,
		AudioTranscriptionFormatSrt,
		AudioTranscriptionFormatText,
		AudioTranscriptionFormatVerboseJSON,
		AudioTranscriptionFormatVtt,
	}
}

// AudioTranslationFormat - Defines available options for the underlying response format of output translation information.
type AudioTranslationFormat string

const (
	// AudioTranslationFormatJSON - Use a response body that is a JSON object containing a single 'text' field for the translation.
	AudioTranslationFormatJSON AudioTranslationFormat = "json"
	// AudioTranslationFormatSrt - Use a response body that is plain text in SubRip (SRT) format that also includes timing information.
	AudioTranslationFormatSrt AudioTranslationFormat = "srt"
	// AudioTranslationFormatText - Use a response body that is plain text containing the raw, unannotated translation.
	AudioTranslationFormatText AudioTranslationFormat = "text"
	// AudioTranslationFormatVerboseJSON - Use a response body that is a JSON object containing translation text along with timing,
	// segments, and other
	// metadata.
	AudioTranslationFormatVerboseJSON AudioTranslationFormat = "verbose_json"
	// AudioTranslationFormatVtt - Use a response body that is plain text in Web Video Text Tracks (VTT) format that also includes
	// timing information.
	AudioTranslationFormatVtt AudioTranslationFormat = "vtt"
)

// PossibleAudioTranslationFormatValues returns the possible values for the AudioTranslationFormat const type.
func PossibleAudioTranslationFormatValues() []AudioTranslationFormat {
	return []AudioTranslationFormat{
		AudioTranslationFormatJSON,
		AudioTranslationFormatSrt,
		AudioTranslationFormatText,
		AudioTranslationFormatVerboseJSON,
		AudioTranslationFormatVtt,
	}
}

// AzureChatExtensionType - A representation of configuration data for a single Azure OpenAI chat extension. This will be
// used by a chat completions request that should use Azure OpenAI chat extensions to augment the response
// behavior. The use of this configuration is compatible only with Azure OpenAI.
type AzureChatExtensionType string

const (
	// AzureChatExtensionTypeAzureCognitiveSearch - Represents the use of Azure Cognitive Search as an Azure OpenAI chat extension.
	AzureChatExtensionTypeAzureCognitiveSearch AzureChatExtensionType = "AzureCognitiveSearch"
	// AzureChatExtensionTypeAzureCosmosDB - Represents the use of Azure Cosmos DB as an Azure OpenAI chat extension.
	AzureChatExtensionTypeAzureCosmosDB AzureChatExtensionType = "AzureCosmosDB"
	// AzureChatExtensionTypeAzureMachineLearningIndex - Represents the use of Azure Machine Learning index as an Azure OpenAI
	// chat extension.
	AzureChatExtensionTypeAzureMachineLearningIndex AzureChatExtensionType = "AzureMLIndex"
	// AzureChatExtensionTypeElasticsearch - Represents the use of Elasticsearch® index as an Azure OpenAI chat extension.
	AzureChatExtensionTypeElasticsearch AzureChatExtensionType = "Elasticsearch"
	// AzureChatExtensionTypePinecone - Represents the use of Pinecone index as an Azure OpenAI chat extension.
	AzureChatExtensionTypePinecone AzureChatExtensionType = "Pinecone"
)

// PossibleAzureChatExtensionTypeValues returns the possible values for the AzureChatExtensionType const type.
func PossibleAzureChatExtensionTypeValues() []AzureChatExtensionType {
	return []AzureChatExtensionType{
		AzureChatExtensionTypeAzureCognitiveSearch,
		AzureChatExtensionTypeAzureCosmosDB,
		AzureChatExtensionTypeAzureMachineLearningIndex,
		AzureChatExtensionTypeElasticsearch,
		AzureChatExtensionTypePinecone,
	}
}

// AzureCognitiveSearchQueryType - The type of Azure Cognitive Search retrieval query that should be executed when using it
// as an Azure OpenAI chat extension.
type AzureCognitiveSearchQueryType string

const (
	// AzureCognitiveSearchQueryTypeSemantic - Represents the semantic query parser for advanced semantic modeling.
	AzureCognitiveSearchQueryTypeSemantic AzureCognitiveSearchQueryType = "semantic"
	// AzureCognitiveSearchQueryTypeSimple - Represents the default, simple query parser.
	AzureCognitiveSearchQueryTypeSimple AzureCognitiveSearchQueryType = "simple"
	// AzureCognitiveSearchQueryTypeVector - Represents vector search over computed data.
	AzureCognitiveSearchQueryTypeVector AzureCognitiveSearchQueryType = "vector"
	// AzureCognitiveSearchQueryTypeVectorSemanticHybrid - Represents a combination of semantic search and vector data querying.
	AzureCognitiveSearchQueryTypeVectorSemanticHybrid AzureCognitiveSearchQueryType = "vectorSemanticHybrid"
	// AzureCognitiveSearchQueryTypeVectorSimpleHybrid - Represents a combination of the simple query strategy with vector data.
	AzureCognitiveSearchQueryTypeVectorSimpleHybrid AzureCognitiveSearchQueryType = "vectorSimpleHybrid"
)

// PossibleAzureCognitiveSearchQueryTypeValues returns the possible values for the AzureCognitiveSearchQueryType const type.
func PossibleAzureCognitiveSearchQueryTypeValues() []AzureCognitiveSearchQueryType {
	return []AzureCognitiveSearchQueryType{
		AzureCognitiveSearchQueryTypeSemantic,
		AzureCognitiveSearchQueryTypeSimple,
		AzureCognitiveSearchQueryTypeVector,
		AzureCognitiveSearchQueryTypeVectorSemanticHybrid,
		AzureCognitiveSearchQueryTypeVectorSimpleHybrid,
	}
}

// ChatCompletionRequestMessageContentPartImageURLDetail - Specifies the detail level of the image. Learn more in the Vision
// guide [/docs/guides/vision/low-or-high-fidelity-image-understanding].
type ChatCompletionRequestMessageContentPartImageURLDetail string

const (
	ChatCompletionRequestMessageContentPartImageURLDetailAuto ChatCompletionRequestMessageContentPartImageURLDetail = "auto"
	ChatCompletionRequestMessageContentPartImageURLDetailHigh ChatCompletionRequestMessageContentPartImageURLDetail = "high"
	ChatCompletionRequestMessageContentPartImageURLDetailLow  ChatCompletionRequestMessageContentPartImageURLDetail = "low"
)

// PossibleChatCompletionRequestMessageContentPartImageURLDetailValues returns the possible values for the ChatCompletionRequestMessageContentPartImageURLDetail const type.
func PossibleChatCompletionRequestMessageContentPartImageURLDetailValues() []ChatCompletionRequestMessageContentPartImageURLDetail {
	return []ChatCompletionRequestMessageContentPartImageURLDetail{
		ChatCompletionRequestMessageContentPartImageURLDetailAuto,
		ChatCompletionRequestMessageContentPartImageURLDetailHigh,
		ChatCompletionRequestMessageContentPartImageURLDetailLow,
	}
}

// ChatCompletionRequestMessageContentPartType - The type of the content part.
type ChatCompletionRequestMessageContentPartType string

const (
	// ChatCompletionRequestMessageContentPartTypeImageURL - Chat content contains an image URL
	ChatCompletionRequestMessageContentPartTypeImageURL ChatCompletionRequestMessageContentPartType = "image_url"
	// ChatCompletionRequestMessageContentPartTypeText - Chat content contains text
	ChatCompletionRequestMessageContentPartTypeText ChatCompletionRequestMessageContentPartType = "text"
)

// PossibleChatCompletionRequestMessageContentPartTypeValues returns the possible values for the ChatCompletionRequestMessageContentPartType const type.
func PossibleChatCompletionRequestMessageContentPartTypeValues() []ChatCompletionRequestMessageContentPartType {
	return []ChatCompletionRequestMessageContentPartType{
		ChatCompletionRequestMessageContentPartTypeImageURL,
		ChatCompletionRequestMessageContentPartTypeText,
	}
}

// ChatRole - A description of the intended purpose of a message within a chat completions interaction.
type ChatRole string

const (
	// ChatRoleAssistant - The role that provides responses to system-instructed, user-prompted input.
	ChatRoleAssistant ChatRole = "assistant"
	// ChatRoleFunction - The role that provides function results for chat completions.
	ChatRoleFunction ChatRole = "function"
	// ChatRoleSystem - The role that instructs or sets the behavior of the assistant.
	ChatRoleSystem ChatRole = "system"
	// ChatRoleTool - The role that represents extension tool activity within a chat completions operation.
	ChatRoleTool ChatRole = "tool"
	// ChatRoleUser - The role that provides input for chat completions.
	ChatRoleUser ChatRole = "user"
)

// PossibleChatRoleValues returns the possible values for the ChatRole const type.
func PossibleChatRoleValues() []ChatRole {
	return []ChatRole{
		ChatRoleAssistant,
		ChatRoleFunction,
		ChatRoleSystem,
		ChatRoleTool,
		ChatRoleUser,
	}
}

// CompletionsFinishReason - Representation of the manner in which a completions response concluded.
type CompletionsFinishReason string

const (
	// CompletionsFinishReasonContentFiltered - Completions generated a response that was identified as potentially sensitive
	// per content
	// moderation policies.
	CompletionsFinishReasonContentFiltered CompletionsFinishReason = "content_filter"
	// CompletionsFinishReasonFunctionCall - Completion ended normally, with the model requesting a function to be called.
	CompletionsFinishReasonFunctionCall CompletionsFinishReason = "function_call"
	// CompletionsFinishReasonStopped - Completions ended normally and reached its end of token generation.
	CompletionsFinishReasonStopped CompletionsFinishReason = "stop"
	// CompletionsFinishReasonTokenLimitReached - Completions exhausted available token limits before generation could complete.
	CompletionsFinishReasonTokenLimitReached CompletionsFinishReason = "length"
	// CompletionsFinishReasonToolCalls - Completion ended with the model calling a provided tool for output.
	CompletionsFinishReasonToolCalls CompletionsFinishReason = "tool_calls"
)

// PossibleCompletionsFinishReasonValues returns the possible values for the CompletionsFinishReason const type.
func PossibleCompletionsFinishReasonValues() []CompletionsFinishReason {
	return []CompletionsFinishReason{
		CompletionsFinishReasonContentFiltered,
		CompletionsFinishReasonFunctionCall,
		CompletionsFinishReasonStopped,
		CompletionsFinishReasonTokenLimitReached,
		CompletionsFinishReasonToolCalls,
	}
}

// ContentFilterSeverity - Ratings for the intensity and risk level of harmful content.
type ContentFilterSeverity string

const (
	// ContentFilterSeverityHigh - Content that displays explicit and severe harmful instructions, actions,
	// damage, or abuse; includes endorsement, glorification, or promotion of severe
	// harmful acts, extreme or illegal forms of harm, radicalization, or non-consensual
	// power exchange or abuse.
	ContentFilterSeverityHigh ContentFilterSeverity = "high"
	// ContentFilterSeverityLow - Content that expresses prejudiced, judgmental, or opinionated views, includes offensive
	// use of language, stereotyping, use cases exploring a fictional world (for example, gaming,
	// literature) and depictions at low intensity.
	ContentFilterSeverityLow ContentFilterSeverity = "low"
	// ContentFilterSeverityMedium - Content that uses offensive, insulting, mocking, intimidating, or demeaning language
	// towards specific identity groups, includes depictions of seeking and executing harmful
	// instructions, fantasies, glorification, promotion of harm at medium intensity.
	ContentFilterSeverityMedium ContentFilterSeverity = "medium"
	// ContentFilterSeveritySafe - Content may be related to violence, self-harm, sexual, or hate categories but the terms
	// are used in general, journalistic, scientific, medical, and similar professional contexts,
	// which are appropriate for most audiences.
	ContentFilterSeveritySafe ContentFilterSeverity = "safe"
)

// PossibleContentFilterSeverityValues returns the possible values for the ContentFilterSeverity const type.
func PossibleContentFilterSeverityValues() []ContentFilterSeverity {
	return []ContentFilterSeverity{
		ContentFilterSeverityHigh,
		ContentFilterSeverityLow,
		ContentFilterSeverityMedium,
		ContentFilterSeveritySafe,
	}
}

// ElasticsearchQueryType - The type of Elasticsearch® retrieval query that should be executed when using it as an Azure OpenAI
// chat extension.
type ElasticsearchQueryType string

const (
	// ElasticsearchQueryTypeSimple - Represents the default, simple query parser.
	ElasticsearchQueryTypeSimple ElasticsearchQueryType = "simple"
	// ElasticsearchQueryTypeVector - Represents vector search over computed data.
	ElasticsearchQueryTypeVector ElasticsearchQueryType = "vector"
)

// PossibleElasticsearchQueryTypeValues returns the possible values for the ElasticsearchQueryType const type.
func PossibleElasticsearchQueryTypeValues() []ElasticsearchQueryType {
	return []ElasticsearchQueryType{
		ElasticsearchQueryTypeSimple,
		ElasticsearchQueryTypeVector,
	}
}

// ImageGenerationQuality - An image generation configuration that specifies how the model should prioritize quality, cost,
// and speed. Only configurable with dall-e-3 models.
type ImageGenerationQuality string

const (
	// ImageGenerationQualityHd - Requests image generation with higher quality, higher cost and lower speed relative to standard.
	ImageGenerationQualityHd ImageGenerationQuality = "hd"
	// ImageGenerationQualityStandard - Requests image generation with standard, balanced characteristics of quality, cost, and
	// speed.
	ImageGenerationQualityStandard ImageGenerationQuality = "standard"
)

// PossibleImageGenerationQualityValues returns the possible values for the ImageGenerationQuality const type.
func PossibleImageGenerationQualityValues() []ImageGenerationQuality {
	return []ImageGenerationQuality{
		ImageGenerationQualityHd,
		ImageGenerationQualityStandard,
	}
}

// ImageGenerationResponseFormat - The format in which the generated images are returned.
type ImageGenerationResponseFormat string

const (
	// ImageGenerationResponseFormatBase64 - Image generation response items should provide image data as a base64-encoded string.
	ImageGenerationResponseFormatBase64 ImageGenerationResponseFormat = "b64_json"
	// ImageGenerationResponseFormatURL - Image generation response items should provide a URL from which the image may be retrieved.
	ImageGenerationResponseFormatURL ImageGenerationResponseFormat = "url"
)

// PossibleImageGenerationResponseFormatValues returns the possible values for the ImageGenerationResponseFormat const type.
func PossibleImageGenerationResponseFormatValues() []ImageGenerationResponseFormat {
	return []ImageGenerationResponseFormat{
		ImageGenerationResponseFormatBase64,
		ImageGenerationResponseFormatURL,
	}
}

// ImageGenerationStyle - An image generation configuration that specifies how the model should incorporate realism and other
// visual characteristics. Only configurable with dall-e-3 models.
type ImageGenerationStyle string

const (
	// ImageGenerationStyleNatural - Requests image generation in a natural style with less preference for dramatic and hyper-realistic
	// characteristics.
	ImageGenerationStyleNatural ImageGenerationStyle = "natural"
	// ImageGenerationStyleVivid - Requests image generation in a vivid style with a higher preference for dramatic and hyper-realistic
	// characteristics.
	ImageGenerationStyleVivid ImageGenerationStyle = "vivid"
)

// PossibleImageGenerationStyleValues returns the possible values for the ImageGenerationStyle const type.
func PossibleImageGenerationStyleValues() []ImageGenerationStyle {
	return []ImageGenerationStyle{
		ImageGenerationStyleNatural,
		ImageGenerationStyleVivid,
	}
}

// ImageSize - The desired size of generated images.
type ImageSize string

const (
	// ImageSizeSize1024X1024 - A standard, square image size of 1024x1024 pixels.
	// Supported by both dall-e-2 and dall-e-3 models.
	ImageSizeSize1024X1024 ImageSize = "1024x1024"
	// ImageSizeSize1024X1792 - A taller image size of 1792x1024 pixels.
	// Only supported with dall-e-3 models.
	ImageSizeSize1024X1792 ImageSize = "1024x1792"
	// ImageSizeSize1792X1024 - A wider image size of 1024x1792 pixels.
	// Only supported with dall-e-3 models.
	ImageSizeSize1792X1024 ImageSize = "1792x1024"
	// ImageSizeSize256X256 - Very small image size of 256x256 pixels.
	// Only supported with dall-e-2 models.
	ImageSizeSize256X256 ImageSize = "256x256"
	// ImageSizeSize512X512 - A smaller image size of 512x512 pixels.
	// Only supported with dall-e-2 models.
	ImageSizeSize512X512 ImageSize = "512x512"
)

// PossibleImageSizeValues returns the possible values for the ImageSize const type.
func PossibleImageSizeValues() []ImageSize {
	return []ImageSize{
		ImageSizeSize1024X1024,
		ImageSizeSize1024X1792,
		ImageSizeSize1792X1024,
		ImageSizeSize256X256,
		ImageSizeSize512X512,
	}
}

// OnYourDataAuthenticationType - The authentication types supported with Azure OpenAI On Your Data.
type OnYourDataAuthenticationType string

const (
	// OnYourDataAuthenticationTypeAPIKey - Authentication via API key.
	OnYourDataAuthenticationTypeAPIKey OnYourDataAuthenticationType = "APIKey"
	// OnYourDataAuthenticationTypeConnectionString - Authentication via connection string.
	OnYourDataAuthenticationTypeConnectionString OnYourDataAuthenticationType = "ConnectionString"
	// OnYourDataAuthenticationTypeKeyAndKeyID - Authentication via key and key ID pair.
	OnYourDataAuthenticationTypeKeyAndKeyID OnYourDataAuthenticationType = "KeyAndKeyId"
	// OnYourDataAuthenticationTypeSystemAssignedManagedIdentity - Authentication via system-assigned managed identity.
	OnYourDataAuthenticationTypeSystemAssignedManagedIdentity OnYourDataAuthenticationType = "SystemAssignedManagedIdentity"
	// OnYourDataAuthenticationTypeUserAssignedManagedIdentity - Authentication via user-assigned managed identity.
	OnYourDataAuthenticationTypeUserAssignedManagedIdentity OnYourDataAuthenticationType = "UserAssignedManagedIdentity"
)

// PossibleOnYourDataAuthenticationTypeValues returns the possible values for the OnYourDataAuthenticationType const type.
func PossibleOnYourDataAuthenticationTypeValues() []OnYourDataAuthenticationType {
	return []OnYourDataAuthenticationType{
		OnYourDataAuthenticationTypeAPIKey,
		OnYourDataAuthenticationTypeConnectionString,
		OnYourDataAuthenticationTypeKeyAndKeyID,
		OnYourDataAuthenticationTypeSystemAssignedManagedIdentity,
		OnYourDataAuthenticationTypeUserAssignedManagedIdentity,
	}
}

// OnYourDataVectorizationSourceType - Represents the available sources Azure OpenAI On Your Data can use to configure vectorization
// of data for use with vector search.
type OnYourDataVectorizationSourceType string

const (
	// OnYourDataVectorizationSourceTypeDeploymentName - Represents an Ada model deployment name to use. This model deployment
	// must be in the same Azure OpenAI resource, but
	// On Your Data will use this model deployment via an internal call rather than a public one, which enables vector
	// search even in private networks.
	OnYourDataVectorizationSourceTypeDeploymentName OnYourDataVectorizationSourceType = "DeploymentName"
	// OnYourDataVectorizationSourceTypeEndpoint - Represents vectorization performed by public service calls to an Azure OpenAI
	// embedding model.
	OnYourDataVectorizationSourceTypeEndpoint OnYourDataVectorizationSourceType = "Endpoint"
	// OnYourDataVectorizationSourceTypeModelID - Represents a specific embedding model ID as defined in the search service.
	// Currently only supported by Elasticsearch®.
	OnYourDataVectorizationSourceTypeModelID OnYourDataVectorizationSourceType = "ModelId"
)

// PossibleOnYourDataVectorizationSourceTypeValues returns the possible values for the OnYourDataVectorizationSourceType const type.
func PossibleOnYourDataVectorizationSourceTypeValues() []OnYourDataVectorizationSourceType {
	return []OnYourDataVectorizationSourceType{
		OnYourDataVectorizationSourceTypeDeploymentName,
		OnYourDataVectorizationSourceTypeEndpoint,
		OnYourDataVectorizationSourceTypeModelID,
	}
}
