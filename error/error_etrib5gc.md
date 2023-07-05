### Run gnb
- pran:
```bash
UERANSIM v3.2.6
[2023-07-05 04:38:17.705] [sctp] [info] Trying to establish SCTP connection... (127.0.0.1:38412)
[2023-07-05 04:38:17.737] [sctp] [info] SCTP connection established (127.0.0.1:38412)
[2023-07-05 04:38:17.741] [sctp] [debug] SCTP association setup ascId[11]
[2023-07-05 04:38:17.741] [ngap] [debug] Sending NG Setup Request
[2023-07-05 04:38:17.745] [ngap] [debug] NG Setup Response received
[2023-07-05 04:38:17.745] [ngap] [info] NG Setup procedure is successful
```

### Run ue
- udm:
```bash
INFO[0042] Receive a GenerateAuthData request from AUSF for ue[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0042] Recover SUPI[imsi-208930000000003] from SUCI[suci-0-208-93-0000-0-0-0000000003] is successful  mod="udm:context"
INFO[0042] Authentication vector is generated for ue[SUPI=imsi-208930000000003]  mod=sbi.producer
```

- aufs
```bash
INFO[0035] Receive an UeAuthenticationsPost from AMF for SUCI=suci-0-208-93-0000-0-0-0000000003  mod="ausf:producer"
INFO[0035] Request from UDM an authentication vector for Ue[SUCI=suci-0-208-93-0000-0-0-0000000003] from UDM  mod="ausf:producer"
INFO[0035] Receive from UDM an authentication vector for Ue[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod="ausf:producer"
INFO[0036] Receive a UeAuthenticationsAuthCtxId5gAkaConfirmationPut from AMF for Ue[SUPI=suci-0-208-93-0000-0-0-0000000003]  mod="ausf:producer"
INFO[0036] Ue[SUPI=suci-0-208-93-0000-0-0-0000000003] is authenticated  mod="ausf:producer"
```

- damf: 1 warn in bottom line
```bash
INFO[0033] Receive initial Ue message                    mod=sbi.producer
INFO[0033] Callback is : /ran/208-93/daejeon_4_0         mod=sbi.producer
INFO[0033] service /ran/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0033] add a service instance [id=4] to group /ran/208-93/daejeon  mod=registry
INFO[0033] Send an UeAuthenticationsPost to AUSF [SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0033] Receive authentication information from the AUSF  mod=sbi.producer
INFO[0033] Send a NAS AuthenticationRequest to UE        mod=sbi.producer
INFO[0033] Receive UplinkNasTransport                    mod=sbi.producer
INFO[0033] Send a UeAuthenticationsAuthCtxId5gAkaConfirmationPut to AUSF for UE[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0033] Authentication is success; receive a SUPI=imsi-208930000000003  mod=sbi.producer
INFO[0033] KAMF is generate e7af5492631f0324c462256b0cbe4c63c99df9a855703245bf077d535f2be716  mod=sbi.producer
INFO[0033] Ue [SUPI=imsi-208930000000003] has subscribed Snssai = 1-010203  mod=sbi.producer
INFO[0033] NSSF finds an AMF[amfid=112233] to handle the UE[subscribed slice=1-010203]  mod=sbi.producer
INFO[0033] UeContext 0 is removed from pool              mod=context
INFO[0033] UeContext 0 is terminated                     mod=sbi.producer
WARN[0036] UeContext reaches its end-of-life             mod=sbi.producer
```

- amf:
```bash
INFO[0006] add a service instance [id=5] to group /pcf/208-93  mod=registry
INFO[0031] Receive an InitUeContextRequest from /ran/208-93/daejeon_4_0  mod="amf:sbi.producer"
WARN[0031] No UeContext is found for SUCI[suci-0-208-93-0000-0-0-0000000003], create a new UeContext  mod="amf:gmm"
INFO[0031] service /ran/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0031] add a service instance [id=4] to group /ran/208-93/daejeon  mod=registry
INFO[0031] Start handling of NAS RegistrationRequest     mod="amf:gmm"
INFO[0031] Send a NAS SecurityModeCommand                mod="amf:gmm"
INFO[0031] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0031] Handle Security Mode Complete                 mod="amf:secmode"
INFO[0031] Got an IMEISV for the UE: imeisv-4370816125816151  mod="amf:secmode"
INFO[0031] Security mode establishment completed         mod="amf:gmm"
INFO[0031] Send an AmPolicyAssociation request to PCF    mod="amf:gmm"
INFO[0031] Allocate registration areas is not implemented  mod="amf:gmm"
INFO[0031] Assign ladn information for UE is not implemented  mod="amf:gmm"
INFO[0032] Send a InitialContextSetupRequest to Proxy RAN  mod="amf:gmm"
INFO[0032] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0032] Receive a UplinkNasTransport from Proxy RAN   mod="amf:gmm"
INFO[0032] No SmContext found for PDU session 1          mod="amf:gmm"
INFO[0032] Ue[SUPI=imsi-208930000000003] sends a snssai[{%!s(int32=1) 010203}] for the requested session[1]  mod="amf:context"
INFO[0032] Select Dnn = internet for the SmContext       mod="amf:context"
INFO[0032] Create a SMF consumer /smf/internet/208-93/1-010203  mod="amf:context"
INFO[0032] service /smf/internet/208-93/1-010203 is not in the registry, lets query the controller  mod=registry
INFO[0032] add a service instance [id=6] to group /smf/internet/208-93/1-010203  mod=registry
INFO[0032] Ask Smf to create a smcontext for the request session (Nas message is forwarded too)  mod="amf:gmm"
INFO[0033] Smf returns no error, store the created SmContext  mod="amf:gmm"
INFO[0033] Store SmContext [sid=1] for Ue[SUPI=imsi-208930000000003]  mod="amf:context"
INFO[0034] Receive a N1N2MessageTransfer from SMF [SUPI=imsi-208930000000003]  mod="amf:sbi.producer"
INFO[0034] Found an SmContext[supi=imsi-208930000000003-sid=1] to  handle N1N2message transfer  mod="amf:gmm"
INFO[0034] Receive a N2 SM Message (PDU Session ID: 1)   mod="amf:gmm"
INFO[0034] Send a PduSessionResourceSetupRequest to Proxy RAN  mod="amf:gmm"
```

- pran: 1 warn line 6, 6 error in bottom line
```bash
INFO[0021] SCTP Accept from: 127.0.0.1:36087             tag="pran:ngap"
INFO[0021] Create a new NG connection for: 127.0.0.1:36087  tag="pran:ngap"
INFO[0021] Receive a  NG Setup request                   tag="pran:ngap"
INFO[0021] Send an NG-Setup response                     tag="pran:ngap"
INFO[0026] Receive an Initial UE Message from gnB        tag="pran:ngap"
WARN[0026] Ue for RanNgapId=1 not found                  tag=ran
INFO[0026] Add UeContext [CuNgapId=0]                    tag="pran:context"
INFO[0026] A new UeContext [RanNgapID: 1, CuNgapId: 0] is created  tag="pran:ngap"
INFO[0026] service /damf/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0026] add a service instance [id=2] to group /damf/208-93/daejeon  mod=registry
INFO[0026] Receive a Downlink Nas message from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0026] Send a DownlinkNasTransport to gnB for ue[ranNgapId=1]  tag=ue
INFO[0026] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0026] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0027] Receive a InitUeContextStatus from AMF for UE [cuNgapId=0]  mod="pran:sbi.producer"
INFO[0027] service /amf/208-93/112233 is not in the registry, lets query the controller  mod=registry
INFO[0027] add a service instance [id=3] to group /amf/208-93/112233  mod=registry
INFO[0027] Connected state                               tag=ue
INFO[0028] Receive AmfUeId=0                             tag=ue
INFO[0028] Receive a Downlink Nas message from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0028] Connected state                               tag=ue
INFO[0028] Send a DownlinkNasTransport to gnB for ue[ranNgapId=1]  tag=ue
INFO[0028] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0028] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0028] Connected state                               tag=ue
INFO[0028] Receive a InitCtxSetupReq from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0028] Connected state                               tag=ue
INFO[0028] Send initial context setup with [CoreNgapId=0][RanNgapId=1] to gnB  tag=ue
INFO[0028] Receive an Initial Context Setup Response from gnB  tag="pran:ngap"
INFO[0028] Connected state                               tag=ue
INFO[0028] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0028] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0028] Connected state                               tag=ue
INFO[0030] Receive a PduSessResSetReq from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0030] Connected state                               tag=ue
INFO[0030] Send a PduSEssionResourceSetupRequest to gnB  tag=ue
INFO[0030] Receive a PDU Session Resource Setup Response from gnB   tag="pran:ngap"
INFO[0030] Connected state                               tag=ue
ERRO[0114] AcceptSCTP: interrupted system call           tag="pran:ngap"
ERRO[0215] AcceptSCTP: interrupted system call           tag="pran:ngap"
ERRO[0217] AcceptSCTP: interrupted system call           tag="pran:ngap"
ERRO[0245] AcceptSCTP: interrupted system call           tag="pran:ngap"
ERRO[0369] AcceptSCTP: interrupted system call           tag="pran:ngap"
ERRO[0385] SCTPRead: interrupted system call             tag="pran:ngap"
```

- gnb:
```bash
UERANSIM v3.1.0
[2023-07-04 07:27:17.912] [sctp] [info] Trying to establish SCTP connection... (127.0.0.1:38412)
[2023-07-04 07:27:17.953] [sctp] [info] SCTP connection established (127.0.0.1:38412)
[2023-07-04 07:27:17.961] [sctp] [debug] SCTP association setup ascId[43]
[2023-07-04 07:27:17.961] [ngap] [debug] Sending NG Setup Request
[2023-07-04 07:27:17.981] [ngap] [debug] NG Setup Response received
[2023-07-04 07:27:17.981] [ngap] [info] NG Setup procedure is successful
[2023-07-04 07:27:35.414] [mr] [info] New UE connected to gNB. Total number of UEs is now: 1
[2023-07-04 07:27:35.415] [rrc] [debug] Sending RRC Setup for UE[3]
[2023-07-04 07:27:35.415] [ngap] [debug] Initial NAS message received from UE 3
```

- pcf: 1 warn
```bash
WARN[0025] Receive a CreateIndividualAMPolicyAssociation from AMF for ue[SUPI=imsi-208930000000003], return a dummy PolicyAssociation  mod=sbi.producer
INFO[0027] Receive a CreateSmPolicy request from SMF for pdu session 1 of ue[SUPI=imsi-208930000000003]; return a dummy SmPolicyDecision  mod=sbi.producer
```

- smf:
```bash
INFO[0023] Create a SMContext for ue[SUPI=imsi-208930000000003]  mod="smf:sm"
INFO[0023] Create a PCF consumer [/pcf/208-93] for session 1 [SUPI=imsi-208930000000003]  mod="smf:sm"
INFO[0023] service /amf/208-93/112233 is not in the registry, lets query the controller  mod=registry
INFO[0024] add a service instance [id=3] to group /amf/208-93/112233  mod=registry
INFO[0024] Amf consumer is created: /amf/208-93/112233/3  mod="smf:sm"
INFO[0024] Handle PDUSessionEstablishmentRequest [session=1,SUPI=imsi-208930000000003]  mod="smf:sm"
INFO[0024] Send a CreateSMPolicy request to PCF          mod="smf:sm"
INFO[0024] Start activating the session supi=imsi-208930000000003-sid=1  mod="smf:sm"
INFO[0024] Looking for a data path (UPFs) in 1-010203 slice that reaches to Dnn=internet  mod="smf:sm"
INFO[0024] A data path has been setup for the session supi=imsi-208930000000003-sid=1  mod="smf:sm"
INFO[0024] Establish pfcp session with UPF(s) [Session=1, SUPI=imsi-208930000000003]  mod="smf:sm"
INFO[0024] Send pfcp session establishment to UPF[192.168.56.104]  mod="smf:upman"
INFO[0024] A PduSEssionEstablishment is built for session 1 [SUPI=imsi-208930000000003]  mod="smf:sm"
INFO[0024] Send N1N2 message transfer for imsi-208930000000003 to AMF  mod="smf:sm"
```

- gnb:
```bash
[2023-07-05 05:32:14.277] [rrc] [debug] UE[1] new signal detected
[2023-07-05 05:32:14.349] [rrc] [info] RRC Setup for UE[1]
[2023-07-05 05:32:14.349] [ngap] [debug] Initial NAS message received from UE[1]
[2023-07-05 05:32:16.522] [ngap] [debug] Initial Context Setup Request received
[2023-07-05 05:32:18.898] [ngap] [info] PDU session resource(s) setup for UE[1] count[1]
```

- ue:
```bash
UERANSIM v3.2.6
[2023-07-05 05:32:14.277] [nas] [info] UE switches to state [MM-DEREGISTERED/PLMN-SEARCH]
[2023-07-05 05:32:14.277] [rrc] [debug] New signal detected for cell[1], total [1] cells in coverage
[2023-07-05 05:32:14.277] [nas] [info] Selected plmn[208/93]
[2023-07-05 05:32:14.277] [rrc] [info] Selected cell plmn[208/93] tac[1] category[SUITABLE]
[2023-07-05 05:32:14.277] [nas] [info] UE switches to state [MM-DEREGISTERED/PS]
[2023-07-05 05:32:14.277] [nas] [info] UE switches to state [MM-DEREGISTERED/NORMAL-SERVICE]
[2023-07-05 05:32:14.277] [nas] [debug] Initial registration required due to [MM-DEREG-NORMAL-SERVICE]
[2023-07-05 05:32:14.349] [nas] [debug] UAC access attempt is allowed for identity[0], category[MO_sig]
[2023-07-05 05:32:14.349] [nas] [debug] Sending Initial Registration
[2023-07-05 05:32:14.349] [rrc] [debug] Sending RRC Setup Request
[2023-07-05 05:32:14.349] [nas] [info] UE switches to state [MM-REGISTER-INITIATED]
[2023-07-05 05:32:14.349] [rrc] [info] RRC connection established
[2023-07-05 05:32:14.349] [rrc] [info] UE switches to state [RRC-CONNECTED]
[2023-07-05 05:32:14.349] [nas] [info] UE switches to state [CM-CONNECTED]
[2023-07-05 05:32:14.890] [nas] [debug] Authentication Request received
[2023-07-05 05:32:16.181] [nas] [debug] Security Mode Command received
[2023-07-05 05:32:16.181] [nas] [debug] Performing kAMF derivation from kAMF in mobility
[2023-07-05 05:32:16.181] [nas] [debug] Selected integrity[1] ciphering[1]
[2023-07-05 05:32:16.522] [nas] [debug] Registration accept received
[2023-07-05 05:32:16.522] [nas] [info] UE switches to state [MM-REGISTERED/NORMAL-SERVICE]
[2023-07-05 05:32:16.522] [nas] [info] Initial Registration is successful
[2023-07-05 05:32:16.522] [nas] [debug] Sending PDU Session Establishment Request
[2023-07-05 05:32:16.525] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 05:32:18.899] [nas] [debug] PDU Session Establishment Accept received
[2023-07-05 05:32:18.899] [nas] [info] PDU Session establishment is successful PSI[1]
[2023-07-05 05:32:19.049] [app] [info] Connection setup for PDU session[1] is successful, TUN interface[uesimtun0, 10.10.10.1] is up.
```