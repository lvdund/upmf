package consumer

// func SendNFRegistration() error {
// 	smfProfile := smf_context.NFProfile

// 	sNssais := []models.Snssai{}
// 	for _, snssaiSmfInfo := range *smfProfile.SMFInfo.SNssaiSmfInfoList {
// 		sNssais = append(sNssais, *snssaiSmfInfo.SNssai)
// 	}

// 	// set nfProfile
// 	profile := models.NfProfile{
// 		NfInstanceId:  smf_context.SMF_Self().NfInstanceID,
// 		NfType:        models.NfType_SMF,
// 		NfStatus:      models.NfStatus_REGISTERED,
// 		Ipv4Addresses: []string{smf_context.SMF_Self().RegisterIPv4},
// 		NfServices:    smfProfile.NFServices,
// 		SmfInfo:       smfProfile.SMFInfo,
// 		SNssais:       &sNssais,
// 		PlmnList:      smfProfile.PLMNList,
// 	}
// 	if smf_context.SMF_Self().Locality != "" {
// 		profile.Locality = smf_context.SMF_Self().Locality
// 	}
// 	var rep models.NfProfile
// 	var res *http.Response
// 	var err error

// 	// Check data (Use RESTful PUT)
// 	for {
// 		rep, res, err = smf_context.SMF_Self().
// 			NFManagementClient.
// 			NFInstanceIDDocumentApi.
// 			RegisterNFInstance(context.TODO(), smf_context.SMF_Self().NfInstanceID, profile)
// 		if err != nil || res == nil {
// 			logger.ConsumerLog.Infof("SMF register to NRF Error[%s]", err.Error())
// 			time.Sleep(2 * time.Second)
// 			continue
// 		}
// 		defer func() {
// 			if resCloseErr := res.Body.Close(); resCloseErr != nil {
// 				logger.ConsumerLog.Errorf("RegisterNFInstance response body cannot close: %+v", resCloseErr)
// 			}
// 		}()

// 		status := res.StatusCode
// 		if status == http.StatusOK {
// 			// NFUpdate
// 			break
// 		} else if status == http.StatusCreated {
// 			// NFRegister
// 			resourceUri := res.Header.Get("Location")
// 			// resouceNrfUri := resourceUri[strings.LastIndex(resourceUri, "/"):]
// 			smf_context.SMF_Self().NfInstanceID = resourceUri[strings.LastIndex(resourceUri, "/")+1:]
// 			break
// 		} else {
// 			logger.ConsumerLog.Infof("handler returned wrong status code %d", status)
// 			// fmt.Errorf("NRF return wrong status code %d", status)
// 		}
// 	}

// 	logger.InitLog.Infof("SMF Registration to NRF %v", rep)
// 	return nil
// }