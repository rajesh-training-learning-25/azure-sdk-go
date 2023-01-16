//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

import "encoding/json"

func unmarshalAssetReferenceBaseClassification(rawMsg json.RawMessage) (AssetReferenceBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AssetReferenceBaseClassification
	switch m["referenceType"] {
	case string(ReferenceTypeDataPath):
		b = &DataPathAssetReference{}
	case string(ReferenceTypeID):
		b = &IDAssetReference{}
	case string(ReferenceTypeOutputPath):
		b = &OutputPathAssetReference{}
	default:
		b = &AssetReferenceBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAutoMLVerticalClassification(rawMsg json.RawMessage) (AutoMLVerticalClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AutoMLVerticalClassification
	switch m["taskType"] {
	case string(TaskTypeClassification):
		b = &Classification{}
	case string(TaskTypeForecasting):
		b = &Forecasting{}
	case string(TaskTypeImageClassification):
		b = &ImageClassification{}
	case string(TaskTypeImageClassificationMultilabel):
		b = &ImageClassificationMultilabel{}
	case string(TaskTypeImageInstanceSegmentation):
		b = &ImageInstanceSegmentation{}
	case string(TaskTypeImageObjectDetection):
		b = &ImageObjectDetection{}
	case string(TaskTypeRegression):
		b = &Regression{}
	case string(TaskTypeTextClassification):
		b = &TextClassification{}
	case string(TaskTypeTextClassificationMultilabel):
		b = &TextClassificationMultilabel{}
	case string(TaskTypeTextNER):
		b = &TextNer{}
	default:
		b = &AutoMLVertical{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalComputeClassification(rawMsg json.RawMessage) (ComputeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AKS{}
	case string(ComputeTypeAmlCompute):
		b = &AmlCompute{}
	case string(ComputeTypeComputeInstance):
		b = &ComputeInstance{}
	case string(ComputeTypeDataFactory):
		b = &DataFactory{}
	case string(ComputeTypeDataLakeAnalytics):
		b = &DataLakeAnalytics{}
	case string(ComputeTypeDatabricks):
		b = &Databricks{}
	case string(ComputeTypeHDInsight):
		b = &HDInsight{}
	case string(ComputeTypeKubernetes):
		b = &Kubernetes{}
	case string(ComputeTypeSynapseSpark):
		b = &SynapseSpark{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachine{}
	default:
		b = &Compute{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalComputeSecretsClassification(rawMsg json.RawMessage) (ComputeSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeSecretsClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AksComputeSecrets{}
	case string(ComputeTypeDatabricks):
		b = &DatabricksComputeSecrets{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachineSecrets{}
	default:
		b = &ComputeSecrets{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataVersionBasePropertiesClassification(rawMsg json.RawMessage) (DataVersionBasePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataVersionBasePropertiesClassification
	switch m["dataType"] {
	case string(DataTypeMltable):
		b = &MLTableData{}
	case string(DataTypeURIFile):
		b = &URIFileDataVersion{}
	case string(DataTypeURIFolder):
		b = &URIFolderDataVersion{}
	default:
		b = &DataVersionBaseProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDatastoreCredentialsClassification(rawMsg json.RawMessage) (DatastoreCredentialsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastoreCredentialsClassification
	switch m["credentialsType"] {
	case string(CredentialsTypeAccountKey):
		b = &AccountKeyDatastoreCredentials{}
	case string(CredentialsTypeCertificate):
		b = &CertificateDatastoreCredentials{}
	case string(CredentialsTypeNone):
		b = &NoneDatastoreCredentials{}
	case string(CredentialsTypeSas):
		b = &SasDatastoreCredentials{}
	case string(CredentialsTypeServicePrincipal):
		b = &ServicePrincipalDatastoreCredentials{}
	default:
		b = &DatastoreCredentials{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDatastorePropertiesClassification(rawMsg json.RawMessage) (DatastorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastorePropertiesClassification
	switch m["datastoreType"] {
	case string(DatastoreTypeAzureBlob):
		b = &AzureBlobDatastore{}
	case string(DatastoreTypeAzureDataLakeGen1):
		b = &AzureDataLakeGen1Datastore{}
	case string(DatastoreTypeAzureDataLakeGen2):
		b = &AzureDataLakeGen2Datastore{}
	case string(DatastoreTypeAzureFile):
		b = &AzureFileDatastore{}
	default:
		b = &DatastoreProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDatastoreSecretsClassification(rawMsg json.RawMessage) (DatastoreSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastoreSecretsClassification
	switch m["secretsType"] {
	case string(SecretsTypeAccountKey):
		b = &AccountKeyDatastoreSecrets{}
	case string(SecretsTypeCertificate):
		b = &CertificateDatastoreSecrets{}
	case string(SecretsTypeSas):
		b = &SasDatastoreSecrets{}
	case string(SecretsTypeServicePrincipal):
		b = &ServicePrincipalDatastoreSecrets{}
	default:
		b = &DatastoreSecrets{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDistributionConfigurationClassification(rawMsg json.RawMessage) (DistributionConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DistributionConfigurationClassification
	switch m["distributionType"] {
	case string(DistributionTypeMpi):
		b = &Mpi{}
	case string(DistributionTypePyTorch):
		b = &PyTorch{}
	case string(DistributionTypeTensorFlow):
		b = &TensorFlow{}
	default:
		b = &DistributionConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEarlyTerminationPolicyClassification(rawMsg json.RawMessage) (EarlyTerminationPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EarlyTerminationPolicyClassification
	switch m["policyType"] {
	case string(EarlyTerminationPolicyTypeBandit):
		b = &BanditPolicy{}
	case string(EarlyTerminationPolicyTypeMedianStopping):
		b = &MedianStoppingPolicy{}
	case string(EarlyTerminationPolicyTypeTruncationSelection):
		b = &TruncationSelectionPolicy{}
	default:
		b = &EarlyTerminationPolicy{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalForecastHorizonClassification(rawMsg json.RawMessage) (ForecastHorizonClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ForecastHorizonClassification
	switch m["mode"] {
	case string(ForecastHorizonModeAuto):
		b = &AutoForecastHorizon{}
	case string(ForecastHorizonModeCustom):
		b = &CustomForecastHorizon{}
	default:
		b = &ForecastHorizon{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalIdentityConfigurationClassification(rawMsg json.RawMessage) (IdentityConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b IdentityConfigurationClassification
	switch m["identityType"] {
	case string(IdentityConfigurationTypeAMLToken):
		b = &AmlToken{}
	case string(IdentityConfigurationTypeManaged):
		b = &ManagedIdentity{}
	case string(IdentityConfigurationTypeUserIdentity):
		b = &UserIdentity{}
	default:
		b = &IdentityConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobBasePropertiesClassification(rawMsg json.RawMessage) (JobBasePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobBasePropertiesClassification
	switch m["jobType"] {
	case string(JobTypeAutoML):
		b = &AutoMLJob{}
	case string(JobTypeCommand):
		b = &CommandJob{}
	case string(JobTypePipeline):
		b = &PipelineJob{}
	case string(JobTypeSweep):
		b = &SweepJob{}
	default:
		b = &JobBaseProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobInputClassification(rawMsg json.RawMessage) (JobInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobInputClassification
	switch m["jobInputType"] {
	case string(JobInputTypeCustomModel):
		b = &CustomModelJobInput{}
	case string(JobInputTypeLiteral):
		b = &LiteralJobInput{}
	case string(JobInputTypeMlflowModel):
		b = &MLFlowModelJobInput{}
	case string(JobInputTypeMltable):
		b = &MLTableJobInput{}
	case string(JobInputTypeTritonModel):
		b = &TritonModelJobInput{}
	case string(JobInputTypeURIFile):
		b = &URIFileJobInput{}
	case string(JobInputTypeURIFolder):
		b = &URIFolderJobInput{}
	default:
		b = &JobInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobInputClassificationMap(rawMsg json.RawMessage) (map[string]JobInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobInputClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalJobOutputClassification(rawMsg json.RawMessage) (JobOutputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobOutputClassification
	switch m["jobOutputType"] {
	case string(JobOutputTypeCustomModel):
		b = &CustomModelJobOutput{}
	case string(JobOutputTypeMlflowModel):
		b = &MLFlowModelJobOutput{}
	case string(JobOutputTypeMltable):
		b = &MLTableJobOutput{}
	case string(JobOutputTypeTritonModel):
		b = &TritonModelJobOutput{}
	case string(JobOutputTypeURIFile):
		b = &URIFileJobOutput{}
	case string(JobOutputTypeURIFolder):
		b = &URIFolderJobOutput{}
	default:
		b = &JobOutput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobOutputClassificationMap(rawMsg json.RawMessage) (map[string]JobOutputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobOutputClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalNCrossValidationsClassification(rawMsg json.RawMessage) (NCrossValidationsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NCrossValidationsClassification
	switch m["mode"] {
	case string(NCrossValidationsModeAuto):
		b = &AutoNCrossValidations{}
	case string(NCrossValidationsModeCustom):
		b = &CustomNCrossValidations{}
	default:
		b = &NCrossValidations{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOnlineDeploymentPropertiesClassification(rawMsg json.RawMessage) (OnlineDeploymentPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OnlineDeploymentPropertiesClassification
	switch m["endpointComputeType"] {
	case string(EndpointComputeTypeKubernetes):
		b = &KubernetesOnlineDeployment{}
	case string(EndpointComputeTypeManaged):
		b = &ManagedOnlineDeployment{}
	default:
		b = &OnlineDeploymentProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOnlineScaleSettingsClassification(rawMsg json.RawMessage) (OnlineScaleSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OnlineScaleSettingsClassification
	switch m["scaleType"] {
	case string(ScaleTypeDefault):
		b = &DefaultScaleSettings{}
	case string(ScaleTypeTargetUtilization):
		b = &TargetUtilizationScaleSettings{}
	default:
		b = &OnlineScaleSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSamplingAlgorithmClassification(rawMsg json.RawMessage) (SamplingAlgorithmClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SamplingAlgorithmClassification
	switch m["samplingAlgorithmType"] {
	case string(SamplingAlgorithmTypeBayesian):
		b = &BayesianSamplingAlgorithm{}
	case string(SamplingAlgorithmTypeGrid):
		b = &GridSamplingAlgorithm{}
	case string(SamplingAlgorithmTypeRandom):
		b = &RandomSamplingAlgorithm{}
	default:
		b = &SamplingAlgorithm{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalScheduleActionBaseClassification(rawMsg json.RawMessage) (ScheduleActionBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ScheduleActionBaseClassification
	switch m["actionType"] {
	case string(ScheduleActionTypeCreateJob):
		b = &JobScheduleAction{}
	case string(ScheduleActionTypeInvokeBatchEndpoint):
		b = &EndpointScheduleAction{}
	default:
		b = &ScheduleActionBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSeasonalityClassification(rawMsg json.RawMessage) (SeasonalityClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SeasonalityClassification
	switch m["mode"] {
	case string(SeasonalityModeAuto):
		b = &AutoSeasonality{}
	case string(SeasonalityModeCustom):
		b = &CustomSeasonality{}
	default:
		b = &Seasonality{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTargetLagsClassification(rawMsg json.RawMessage) (TargetLagsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetLagsClassification
	switch m["mode"] {
	case string(TargetLagsModeAuto):
		b = &AutoTargetLags{}
	case string(TargetLagsModeCustom):
		b = &CustomTargetLags{}
	default:
		b = &TargetLags{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTargetRollingWindowSizeClassification(rawMsg json.RawMessage) (TargetRollingWindowSizeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetRollingWindowSizeClassification
	switch m["mode"] {
	case string(TargetRollingWindowSizeModeAuto):
		b = &AutoTargetRollingWindowSize{}
	case string(TargetRollingWindowSizeModeCustom):
		b = &CustomTargetRollingWindowSize{}
	default:
		b = &TargetRollingWindowSize{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTriggerBaseClassification(rawMsg json.RawMessage) (TriggerBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerBaseClassification
	switch m["triggerType"] {
	case string(TriggerTypeCron):
		b = &CronTrigger{}
	case string(TriggerTypeRecurrence):
		b = &RecurrenceTrigger{}
	default:
		b = &TriggerBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalWorkspaceConnectionPropertiesV2Classification(rawMsg json.RawMessage) (WorkspaceConnectionPropertiesV2Classification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkspaceConnectionPropertiesV2Classification
	switch m["authType"] {
	case string(ConnectionAuthTypeManagedIdentity):
		b = &ManagedIdentityAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeNone):
		b = &NoneAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypePAT):
		b = &PATAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeSAS):
		b = &SASAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeUsernamePassword):
		b = &UsernamePasswordAuthTypeWorkspaceConnectionProperties{}
	default:
		b = &WorkspaceConnectionPropertiesV2{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
