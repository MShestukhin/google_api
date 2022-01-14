package handlers

import (
	"context"
	"github.com/opteo/google-ads-go/services"
	"gitlab.somin.ai/analytics/platform/services/google_api/internal/google"
	pb "gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"golang.org/x/oauth2"
	"strconv"
	"strings"
)

type googleApiHandler struct {
	credentials google.GoogleAdsClientParams
}

func (g googleApiHandler) GetAdGroupsCriteria(ctx context.Context, request *pb.GetAdGroupsCriteriaRequest, response *pb.GetAdGroupsCriteriaResponse) error {
	requestGroup := services.SearchGoogleAdsRequest{
		CustomerId: request.CustomerId,
		Query: `SELECT
		  ad_group_criterion.keyword.text,
		  ad_group.name,
		  campaign.name,
		  metrics.impressions,
		  metrics.clicks,
		  metrics.ctr,
		  metrics.average_cpc
		FROM keyword_view
		WHERE segments.date DURING LAST_30_DAYS`,
	}

	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	//Get the results
	responseAdGroupCriterion, err := googleAdsService.Search(client.Context(), &requestGroup)
	if err != nil {
		return err
	}
	for _, row := range responseAdGroupCriterion.Results {
		criterion := row.GetAdGroupCriterion()
		group := row.GetAdGroup()
		campaign := row.GetCampaign()
		metrics := row.GetMetrics()
		srcGroup := &pb.AdGroupsCriteria{
			Keyword:      criterion.GetKeyword().String(),
			AdGroupName:  group.GetName(),
			CampaignName: campaign.GetName(),
			Impressions:  metrics.GetImpressions(),
			Clicks:       metrics.GetClicks(),
			Ctr:          float32(metrics.GetCtr()),
			AverageCpc:   float32(metrics.GetAverageCpc()),
		}
		response.AdGroupsCriteria = append(response.AdGroupsCriteria, srcGroup)
	}
	return nil
}

func (g googleApiHandler) GetCampaignCriteria(ctx context.Context, request *pb.GetCampaignCriteriaRequest, response *pb.GetCampaignCriteriaResponse) error {
	requestGroup := services.SearchGoogleAdsRequest{
		CustomerId: request.CustomerId,
		Query: `SELECT
          campaign.id,
          campaign_criterion.campaign,
          campaign_criterion.criterion_id,
          campaign_criterion.negative,
          campaign_criterion.type,
          campaign_criterion.keyword.text,
          campaign_criterion.keyword.match_type
        FROM campaign_criterion
        WHERE campaign.id =` + strconv.FormatInt(request.CampaignId, 10),
	}

	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	//Get the results
	responseCampaignCriterion, err := googleAdsService.Search(client.Context(), &requestGroup)
	if err != nil {
		return err
	}
	for _, row := range responseCampaignCriterion.Results {
		criterion := row.CampaignCriterion
		srcGroup := &pb.CampaignCriteria{
			CriterionId:  criterion.GetCriterionId(),
			ResourceName: criterion.GetResourceName(),
			Campaign:     criterion.GetCampaign(),
			DisplayName:  criterion.GetDisplayName(),
			BidModifier:  criterion.GetBidModifier(),
			Negative:     criterion.GetNegative(),
		}
		response.CampaignCriteria = append(response.CampaignCriteria, srcGroup)
	}
	return nil
}

func (g googleApiHandler) ExchangeCode(ctx context.Context, request *pb.ExchangeCodeRequest, response *pb.ExchangeCodeResponse) error {
	conf := &oauth2.Config{
		ClientID:     g.credentials.ClientID,
		ClientSecret: g.credentials.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/adwords",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL: g.credentials.DeveloperToken,
		},
	}

	tok, err := conf.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		return err
	}
	response = &pb.ExchangeCodeResponse{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		ExpireIn:     tok.Expiry.Unix(),
		CreatedAt:    tok.Expiry.Unix(),
	}
	return nil
}

func (g googleApiHandler) GetOAuthUrl(ctx context.Context, request *pb.GetOAuthUrlRequest, response *pb.GetOAuthUrlResponse) error {
	conf := &oauth2.Config{
		ClientID:     g.credentials.ClientID,
		ClientSecret: g.credentials.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/adwords",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL: g.credentials.DeveloperToken,
		},
		RedirectURL: request.RedirectUrl,
	}
	response.Url = conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return nil
}

func (g googleApiHandler) RefreshToken(ctx context.Context, request *pb.RefreshTockenRequest, response *pb.RefreshTockenResponse) error {
	g.credentials.RefreshToken = request.RefreshToken
	token, err := google.RefreshToken(&g.credentials)
	if err != nil {
		return err
	}
	response.Token = token
	return nil
}

func (g googleApiHandler) GetCustomers(ctx context.Context, request *pb.GetCustomersRequest, response *pb.GetCustomersResponse) error {
	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	customersSvc := services.NewCustomerServiceClient(client.Conn())
	customers, err := customersSvc.ListAccessibleCustomers(client.Context(), &services.ListAccessibleCustomersRequest{})
	if err != nil {
		return err
	}
	for _, customerStr := range customers.ResourceNames {
		response.CustomerIds = append(response.CustomerIds, strings.Split(customerStr, "/")[1])
	}
	return nil
}

func (g *googleApiHandler) GetCampaigns(ctx context.Context, request *pb.GetCampaignsRequest, response *pb.GetCampaignsResponse) error {
	//// Create a search request
	campaignRequest := services.SearchGoogleAdsRequest{
		CustomerId: request.CustomerId,
		Query: `SELECT campaign.id, 
							campaign.name,
							campaign.tracking_url_template,
							campaign.base_campaign,
							campaign.campaign_budget,
							campaign.start_date,
							campaign.end_date,
							campaign.final_url_suffix,
							campaign.optimization_score,
							campaign.resource_name,
							campaign.status,
							campaign.serving_status,
							campaign.ad_serving_optimization_status,
							campaign.advertising_channel_type,
							campaign.advertising_channel_sub_type,
							campaign.bidding_strategy_type,
							campaign.accessible_bidding_strategy,
							campaign.payment_mode,
							campaign.labels,
							campaign.real_time_bidding_setting.opt_in, 
							campaign.url_custom_parameters, 
							campaign.network_settings.target_content_network, 
							campaign.network_settings.target_google_search, 
							campaign.network_settings.target_partner_search_network, 
							campaign.network_settings.target_search_network, 
							campaign.hotel_setting.hotel_center_id, 
							campaign.dynamic_search_ads_setting.domain_name, 
							campaign.dynamic_search_ads_setting.feeds, 
							campaign.dynamic_search_ads_setting.language_code, 
							campaign.dynamic_search_ads_setting.use_supplied_urls_only, 
							campaign.shopping_setting.campaign_priority, 
							campaign.shopping_setting.enable_local, 	
							campaign.shopping_setting.merchant_id, 
							campaign.shopping_setting.sales_country, 
							campaign.geo_target_type_setting.negative_geo_target_type, 
							campaign.geo_target_type_setting.positive_geo_target_type, 	
							campaign.local_campaign_setting.location_source_type, 
							campaign.experiment_type, 
							campaign.video_brand_safety_suitability, 
							campaign.vanity_pharma.vanity_pharma_display_url_mode, 
							campaign.vanity_pharma.vanity_pharma_text, 
							campaign.selective_optimization.conversion_actions, 
							campaign.optimization_goal_setting.optimization_goal_types, 
							campaign.tracking_setting.tracking_url, 
							campaign.excluded_parent_asset_field_types
							FROM campaign ORDER BY campaign.id`,
	}

	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	campaignResponse, err := googleAdsService.Search(client.Context(), &campaignRequest)

	if err != nil {
		return err
	}
	for _, row := range campaignResponse.Results {
		campaign := row.Campaign
		srcCampaign := pb.UnmarshalCampaign(campaign)
		response.Campaign = append(response.Campaign, srcCampaign)
	}
	return nil
}

func (g googleApiHandler) GetAdGroups(ctx context.Context, request *pb.GetAdGroupsRequest, response *pb.GetAdGroupsResponse) error {
	requestGroup := services.SearchGoogleAdsRequest{
		CustomerId: request.CustomerId,
		Query: `SELECT ad_group.id, 
					ad_group.name,
					ad_group.base_ad_group,
					ad_group.tracking_url_template,
					ad_group.campaign,
					ad_group.cpc_bid_micros,
					ad_group.cpm_bid_micros,
					ad_group.target_cpa_micros,
					ad_group.cpv_bid_micros,
					ad_group.target_cpm_micros,
					ad_group.target_roas,
					ad_group.percent_cpc_bid_micros,
					ad_group.final_url_suffix,
					ad_group.effective_target_cpa_micros,
					ad_group.effective_target_roas,
					ad_group.resource_name,
					ad_group.status,
					ad_group.type,
					ad_group.ad_rotation_mode,
					ad_group.labels,
					ad_group.explorer_auto_optimizer_setting.opt_in, 
					ad_group.display_custom_bid_dimension FROM ad_group`,
	}

	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	//Get the results
	responseGroup, err := googleAdsService.Search(client.Context(), &requestGroup)
	if err != nil {
		return err
	}
	for _, row := range responseGroup.Results {
		group := row.AdGroup
		srcGroup := &pb.AdGroup{
			Id:                       group.GetId(),
			Name:                     group.GetName(),
			BaseAdGroup:              group.GetBaseAdGroup(),
			TrackingUrlTemplate:      group.GetTrackingUrlTemplate(),
			Campaign:                 group.GetCampaign(),
			CpcBidMicros:             group.GetCpcBidMicros(),
			CpmBidMicros:             group.GetCpmBidMicros(),
			TargetCpaMicros:          group.GetTargetCpaMicros(),
			CpvBidMicros:             group.GetCpvBidMicros(),
			TargetCpmMicros:          group.GetTargetCpmMicros(),
			TargetRoas:               float32(group.GetTargetRoas()),
			PercentCpcBidMicros:      group.GetPercentCpcBidMicros(),
			FinalUrlSuffix:           group.GetFinalUrlSuffix(),
			EffectiveTargetCpaMicros: group.GetEffectiveTargetCpaMicros(),
			EffectiveTargetRoas:      float32(group.GetEffectiveTargetRoas()),
			ResourceName:             group.GetResourceName(),
			Labels:                   group.GetLabels(),
		}
		response.AdGroup = append(response.AdGroup, srcGroup)
	}
	return nil
}

func (g googleApiHandler) GetAdGroupsAd(ctx context.Context, request *pb.GetAdGroupsAdRequest, response *pb.GetAdGroupsAdResponse) error {
	requestGroupAd := services.SearchGoogleAdsRequest{
		CustomerId: request.CustomerId,
		Query: `SELECT ad_group_ad.ad.resource_name, 
					ad_group_ad.ad.name, 
					ad_group_ad.ad.id,
					ad_group_ad.ad.tracking_url_template, 
					ad_group.final_url_suffix, 
					ad_group_ad.ad.display_url, 
					ad_group_ad.ad.added_by_google_ads, 	
					ad_group_ad.ad.final_urls, 
					ad_group_ad.ad.final_mobile_urls, 
					ad_group_ad.ad.url_custom_parameters, 
					ad_group.excluded_parent_asset_field_types, 
					ad_group_ad.ad.device_preference,
					ad_group_ad.ad.system_managed_resource_source FROM ad_group_ad`,
	}

	client, err := google.NewClient(request.Token, g.credentials.DeveloperToken)
	if err != nil {
		panic(err)
	}

	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	//Get the results
	responseGroupAd, err := googleAdsService.Search(client.Context(), &requestGroupAd)
	if err != nil {
		return err
	}
	for _, row := range responseGroupAd.Results {
		ads := row.AdGroupAd
		dscAds := ads.GetAd()
		srcAds := &pb.AdGroupAd{
			Id:                  dscAds.GetId(),
			Name:                dscAds.GetName(),
			ResourceName:        dscAds.GetResourceName(),
			DisplayUrl:          dscAds.GetDisplayUrl(),
			TrackingUrlTemplate: dscAds.GetTrackingUrlTemplate(),
			AddedByGoogleAds:    dscAds.GetAddedByGoogleAds(),
			FinalUrls:           dscAds.GetFinalUrls(),
			FinalMobileUrls:     dscAds.GetFinalMobileUrls(),
			FinalUrlSuffix:      dscAds.GetFinalUrlSuffix(),
		}
		response.Ads = append(response.Ads, srcAds)
	}
	return nil
}

func NewGoogleApiHandler(credentials google.GoogleAdsClientParams) pb.GoogleApiHandler {
	return &googleApiHandler{
		credentials,
	}
}
