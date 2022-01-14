package google_api

import (
	"github.com/opteo/google-ads-go/enums"
	"github.com/opteo/google-ads-go/resources"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var campaignStatusMap = map[enums.CampaignStatusEnum_CampaignStatus]CampaignStatusEnum_CampaignStatus{
	enums.CampaignStatusEnum_UNSPECIFIED: CampaignStatusEnum_UNSPECIFIED,
	enums.CampaignStatusEnum_UNKNOWN:     CampaignStatusEnum_UNKNOWN,
	enums.CampaignStatusEnum_ENABLED:     CampaignStatusEnum_ENABLED,
	enums.CampaignStatusEnum_PAUSED:      CampaignStatusEnum_PAUSED,
	enums.CampaignStatusEnum_REMOVED:     CampaignStatusEnum_REMOVED,
}

var campaignServingStatusMap = map[enums.CampaignServingStatusEnum_CampaignServingStatus]CampaignServingStatusEnum_CampaignServingStatus{
	enums.CampaignServingStatusEnum_UNSPECIFIED: CampaignServingStatusEnum_UNSPECIFIED,
	enums.CampaignServingStatusEnum_UNKNOWN:     CampaignServingStatusEnum_UNKNOWN,
	enums.CampaignServingStatusEnum_SERVING:     CampaignServingStatusEnum_SERVING,
	enums.CampaignServingStatusEnum_NONE:        CampaignServingStatusEnum_NONE,
	enums.CampaignServingStatusEnum_ENDED:       CampaignServingStatusEnum_ENDED,
	enums.CampaignServingStatusEnum_PENDING:     CampaignServingStatusEnum_PENDING,
	enums.CampaignServingStatusEnum_SUSPENDED:   CampaignServingStatusEnum_SUSPENDED,
}

var campaignAdServingOptimizationStatusMap = map[enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus]AdServingOptimizationStatus{
	enums.AdServingOptimizationStatusEnum_UNSPECIFIED:         AdServingOptimizationStatus_UNSPECIFIED,
	enums.AdServingOptimizationStatusEnum_UNKNOWN:             AdServingOptimizationStatus_UNKNOWN,
	enums.AdServingOptimizationStatusEnum_OPTIMIZE:            AdServingOptimizationStatus_OPTIMIZE,
	enums.AdServingOptimizationStatusEnum_CONVERSION_OPTIMIZE: AdServingOptimizationStatus_CONVERSION_OPTIMIZE,
	enums.AdServingOptimizationStatusEnum_ROTATE:              AdServingOptimizationStatus_ROTATE,
	enums.AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY: AdServingOptimizationStatus_ROTATE_INDEFINITELY,
	enums.AdServingOptimizationStatusEnum_UNAVAILABLE:         AdServingOptimizationStatus_UNAVAILABLE,
}

var campaignPaymentModeEnumPaymentModeMap = map[enums.PaymentModeEnum_PaymentMode]PaymentModeEnum_PaymentMode{
	enums.PaymentModeEnum_UNSPECIFIED:      PaymentModeEnum_UNSPECIFIED,
	enums.PaymentModeEnum_UNKNOWN:          PaymentModeEnum_UNKNOWN,
	enums.PaymentModeEnum_CLICKS:           PaymentModeEnum_CLICKS,
	enums.PaymentModeEnum_CONVERSION_VALUE: PaymentModeEnum_CONVERSION_VALUE,
	enums.PaymentModeEnum_CONVERSIONS:      PaymentModeEnum_CONVERSIONS,
	enums.PaymentModeEnum_GUEST_STAY:       PaymentModeEnum_GUEST_STAY,
}

var campaignNegativeGeoTargetTypMap = map[enums.NegativeGeoTargetTypeEnum_NegativeGeoTargetType]NegativeGeoTargetTypeEnum_NegativeGeoTargetType{
	enums.NegativeGeoTargetTypeEnum_UNSPECIFIED:          NegativeGeoTargetTypeEnum_UNSPECIFIED,
	enums.NegativeGeoTargetTypeEnum_UNKNOWN:              NegativeGeoTargetTypeEnum_UNKNOWN,
	enums.NegativeGeoTargetTypeEnum_PRESENCE_OR_INTEREST: NegativeGeoTargetTypeEnum_PRESENCE_OR_INTEREST,
	enums.NegativeGeoTargetTypeEnum_PRESENCE:             NegativeGeoTargetTypeEnum_PRESENCE,
}

var campaignPositiveGeoTargetTypMap = map[enums.PositiveGeoTargetTypeEnum_PositiveGeoTargetType]PositiveGeoTargetTypeEnum_PositiveGeoTargetType{
	enums.PositiveGeoTargetTypeEnum_UNSPECIFIED:          PositiveGeoTargetTypeEnum_UNSPECIFIED,
	enums.PositiveGeoTargetTypeEnum_UNKNOWN:              PositiveGeoTargetTypeEnum_UNKNOWN,
	enums.PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST: PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST,
	enums.PositiveGeoTargetTypeEnum_SEARCH_INTEREST:      PositiveGeoTargetTypeEnum_SEARCH_INTEREST,
	enums.PositiveGeoTargetTypeEnum_PRESENCE:             PositiveGeoTargetTypeEnum_PRESENCE,
}

var campaignLocationSourceTypeMap = map[enums.LocationSourceTypeEnum_LocationSourceType]LocationSourceTypeEnum_LocationSourceType{
	enums.LocationSourceTypeEnum_UNSPECIFIED:        LocationSourceTypeEnum_UNSPECIFIED,
	enums.LocationSourceTypeEnum_UNKNOWN:            LocationSourceTypeEnum_UNKNOWN,
	enums.LocationSourceTypeEnum_GOOGLE_MY_BUSINESS: LocationSourceTypeEnum_GOOGLE_MY_BUSINESS,
	enums.LocationSourceTypeEnum_AFFILIATE:          LocationSourceTypeEnum_AFFILIATE,
}

var campaignExperimentTypeMap = map[enums.CampaignExperimentTypeEnum_CampaignExperimentType]CampaignExperimentTypeEnum_CampaignExperimentType{
	enums.CampaignExperimentTypeEnum_UNSPECIFIED: CampaignExperimentTypeEnum_UNSPECIFIED,
	enums.CampaignExperimentTypeEnum_UNKNOWN:     CampaignExperimentTypeEnum_UNKNOWN,
	enums.CampaignExperimentTypeEnum_BASE:        CampaignExperimentTypeEnum_BASE,
	enums.CampaignExperimentTypeEnum_DRAFT:       CampaignExperimentTypeEnum_DRAFT,
	enums.CampaignExperimentTypeEnum_EXPERIMENT:  CampaignExperimentTypeEnum_EXPERIMENT,
}

var campaignBrandSafetySuitabilityMap = map[enums.BrandSafetySuitabilityEnum_BrandSafetySuitability]BrandSafetySuitabilityEnum_BrandSafetySuitability{
	enums.BrandSafetySuitabilityEnum_UNSPECIFIED:        BrandSafetySuitabilityEnum_UNSPECIFIED,
	enums.BrandSafetySuitabilityEnum_UNKNOWN:            BrandSafetySuitabilityEnum_UNKNOWN,
	enums.BrandSafetySuitabilityEnum_EXPANDED_INVENTORY: BrandSafetySuitabilityEnum_EXPANDED_INVENTORY,
	enums.BrandSafetySuitabilityEnum_STANDARD_INVENTORY: BrandSafetySuitabilityEnum_STANDARD_INVENTORY,
	enums.BrandSafetySuitabilityEnum_LIMITED_INVENTORY:  BrandSafetySuitabilityEnum_LIMITED_INVENTORY,
}

var campaignVanityPharmaDisplayUrlModeMap = map[enums.VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode]VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode{
	enums.VanityPharmaDisplayUrlModeEnum_UNSPECIFIED:              VanityPharmaDisplayUrlModeEnum_UNSPECIFIED,
	enums.VanityPharmaDisplayUrlModeEnum_UNKNOWN:                  VanityPharmaDisplayUrlModeEnum_UNKNOWN,
	enums.VanityPharmaDisplayUrlModeEnum_MANUFACTURER_WEBSITE_URL: VanityPharmaDisplayUrlModeEnum_MANUFACTURER_WEBSITE_URL,
	enums.VanityPharmaDisplayUrlModeEnum_WEBSITE_DESCRIPTION:      VanityPharmaDisplayUrlModeEnum_WEBSITE_DESCRIPTION,
}

var campaignVanityPharmaTextMap = map[enums.VanityPharmaTextEnum_VanityPharmaText]VanityPharmaTextEnum_VanityPharmaText{
	enums.VanityPharmaTextEnum_UNSPECIFIED:                           VanityPharmaTextEnum_UNSPECIFIED,
	enums.VanityPharmaTextEnum_UNKNOWN:                               VanityPharmaTextEnum_UNKNOWN,
	enums.VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_EN:     VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_EN,
	enums.VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_ES:     VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_ES,
	enums.VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_EN:        VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_EN,
	enums.VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_ES:        VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_ES,
	enums.VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_EN:             VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_EN,
	enums.VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_ES:             VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_ES,
	enums.VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_EN:     VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_EN,
	enums.VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_ES:     VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_ES,
	enums.VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_EN: VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_EN,
	enums.VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_ES: VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_ES,
	enums.VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_EN:       VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_EN,
	enums.VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_ES:       VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_ES,
}

var campaignAssetFieldTypeMap = map[enums.AssetFieldTypeEnum_AssetFieldType]AssetFieldTypeEnum_AssetFieldType{
	enums.AssetFieldTypeEnum_UNSPECIFIED:        AssetFieldTypeEnum_UNSPECIFIED,
	enums.AssetFieldTypeEnum_UNKNOWN:            AssetFieldTypeEnum_UNKNOWN,
	enums.AssetFieldTypeEnum_HEADLINE:           AssetFieldTypeEnum_HEADLINE,
	enums.AssetFieldTypeEnum_DESCRIPTION:        AssetFieldTypeEnum_DESCRIPTION,
	enums.AssetFieldTypeEnum_MANDATORY_AD_TEXT:  AssetFieldTypeEnum_MANDATORY_AD_TEXT,
	enums.AssetFieldTypeEnum_MARKETING_IMAGE:    AssetFieldTypeEnum_MARKETING_IMAGE,
	enums.AssetFieldTypeEnum_MEDIA_BUNDLE:       AssetFieldTypeEnum_MEDIA_BUNDLE,
	enums.AssetFieldTypeEnum_YOUTUBE_VIDEO:      AssetFieldTypeEnum_YOUTUBE_VIDEO,
	enums.AssetFieldTypeEnum_BOOK_ON_GOOGLE:     AssetFieldTypeEnum_BOOK_ON_GOOGLE,
	enums.AssetFieldTypeEnum_LEAD_FORM:          AssetFieldTypeEnum_LEAD_FORM,
	enums.AssetFieldTypeEnum_PROMOTION:          AssetFieldTypeEnum_PROMOTION,
	enums.AssetFieldTypeEnum_CALLOUT:            AssetFieldTypeEnum_CALLOUT,
	enums.AssetFieldTypeEnum_STRUCTURED_SNIPPET: AssetFieldTypeEnum_STRUCTURED_SNIPPET,
	enums.AssetFieldTypeEnum_SITELINK:           AssetFieldTypeEnum_SITELINK,
}

func UnmarshalCampaign(campaign *resources.Campaign) *Campaign {

	return &Campaign{
		state:                       protoimpl.MessageState{},
		Id:                          campaign.GetId(),
		Name:                        campaign.GetName(),
		TrackingUrlTemplate:         campaign.GetTrackingUrlTemplate(),
		BaseCampaign:                campaign.GetBaseCampaign(),
		CampaignBudget:              campaign.GetCampaignBudget(),
		StartDate:                   campaign.GetStartDate(),
		EndDate:                     campaign.GetEndDate(),
		FinalUrlSuffix:              campaign.GetFinalUrlSuffix(),
		OptimizationScore:           float32(campaign.GetOptimizationScore()),
		ResourceName:                campaign.GetResourceName(),
		Status:                      campaignStatusMap[campaign.GetStatus()],
		ServingStatus:               campaignServingStatusMap[campaign.GetServingStatus()],
		AdServingOptimizationStatus: campaignAdServingOptimizationStatusMap[campaign.GetAdServingOptimizationStatus()],
		AccessibleBiddingStrategy:   campaign.GetAccessibleBiddingStrategy(),
		PaymentMode:                 campaignPaymentModeEnumPaymentModeMap[campaign.GetPaymentMode()],
		Labels:                      campaign.GetLabels(),
		RealTimeBiddingSetting:      campaign.GetRealTimeBiddingSetting().GetOptIn(),
		NetworkSettings: &NetworkSetting{
			TargetContentNetwork:       campaign.GetNetworkSettings().GetTargetContentNetwork(),
			TargetGoogleSearch:         campaign.GetNetworkSettings().GetTargetGoogleSearch(),
			TargetPartnerSearchNetwork: campaign.GetNetworkSettings().GetTargetPartnerSearchNetwork(),
			TargetSearchNetwork:        campaign.GetNetworkSettings().GetTargetSearchNetwork(),
		},
		HottelSettings: &HotelSetting{
			HotelCenterId: campaign.GetHotelSetting().GetHotelCenterId(),
		},
		DynamicSearchAdsSetting: &DynamicSearchAdsSetting{
			DomainName:          campaign.GetDynamicSearchAdsSetting().GetDomainName(),
			Feeds:               campaign.GetDynamicSearchAdsSetting().GetFeeds(),
			LanguageCode:        campaign.GetDynamicSearchAdsSetting().GetLanguageCode(),
			UseSuppliedUrlsOnly: campaign.GetDynamicSearchAdsSetting().GetUseSuppliedUrlsOnly(),
		},
		ShoppingSetting: &ShoppingSetting{
			CampaignPriority: campaign.GetShoppingSetting().GetCampaignPriority(),
			EnableLocal:      campaign.GetShoppingSetting().GetEnableLocal(),
			MerchantId:       campaign.GetShoppingSetting().GetMerchantId(),
			SalesCountry:     campaign.GetShoppingSetting().GetSalesCountry(),
		},
		GeoTargetTypeSetting: &GeoTargetTypeSetting{
			NegativeGeoTargetType: campaignNegativeGeoTargetTypMap[campaign.GetGeoTargetTypeSetting().GetNegativeGeoTargetType()],
			PositiveGeoTargetType: campaignPositiveGeoTargetTypMap[campaign.GetGeoTargetTypeSetting().GetPositiveGeoTargetType()],
		},
		LocalCampaignSetting:        campaignLocationSourceTypeMap[campaign.GetLocalCampaignSetting().GetLocationSourceType()],
		ExperimentType:              campaignExperimentTypeMap[campaign.GetExperimentType()],
		VideoBrandSafetySuitability: campaignBrandSafetySuitabilityMap[campaign.GetVideoBrandSafetySuitability()],
		VanityPharmaDisplayUrlMode:  campaignVanityPharmaDisplayUrlModeMap[campaign.GetVanityPharma().GetVanityPharmaDisplayUrlMode()],
		VanityPharmaText:            campaignVanityPharmaTextMap[campaign.GetVanityPharma().GetVanityPharmaText()],
		SelectiveOptimization: &SelectiveOptimization{
			ConversionActions: campaign.GetSelectiveOptimization().GetConversionActions(),
		},
		TrackingSetting: &TrackingSetting{
			TrackingUrl: campaign.GetTrackingSetting().GetTrackingUrl(),
		},
	}
}
