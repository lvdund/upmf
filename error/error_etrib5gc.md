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
INFO[0198] Receive a GenerateAuthData request from AUSF for ue[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0198] Recover SUPI[imsi-208930000000003] from SUCI[suci-0-208-93-0000-0-0-0000000003] is successful  mod="udm:context"
INFO[0198] Authentication vector is generated for ue[SUPI=imsi-208930000000003]  mod=sbi.producer
```

- aufs
```bash
INFO[0169] Receive an UeAuthenticationsPost from AMF for SUCI=suci-0-208-93-0000-0-0-0000000003  mod="ausf:producer"
INFO[0169] Request from UDM an authentication vector for Ue[SUCI=suci-0-208-93-0000-0-0-0000000003] from UDM  mod="ausf:producer"
INFO[0169] Receive from UDM an authentication vector for Ue[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod="ausf:producer"
INFO[0169] Receive a UeAuthenticationsAuthCtxId5gAkaConfirmationPut from AMF for Ue[SUPI=suci-0-208-93-0000-0-0-0000000003]  mod="ausf:producer"
INFO[0169] Ue[SUPI=suci-0-208-93-0000-0-0-0000000003] is authenticated  mod="ausf:producer"
```

- damf:
```bash
INFO[0146] Receive initial Ue message                    mod=sbi.producer
INFO[0146] Callback is : /ran/208-93/daejeon_4_0         mod=sbi.producer
INFO[0146] service /ran/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0146] add a service instance [id=4] to group /ran/208-93/daejeon  mod=registry
INFO[0146] Send an UeAuthenticationsPost to AUSF [SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0146] Receive authentication information from the AUSF  mod=sbi.producer
INFO[0146] Send a NAS AuthenticationRequest to UE        mod=sbi.producer
INFO[0146] Receive UplinkNasTransport                    mod=sbi.producer
INFO[0146] Send a UeAuthenticationsAuthCtxId5gAkaConfirmationPut to AUSF for UE[SUCI=suci-0-208-93-0000-0-0-0000000003]  mod=sbi.producer
INFO[0146] Authentication is success; receive a SUPI=imsi-208930000000003  mod=sbi.producer
INFO[0146] KAMF is generate 09af428ef6d7cfaaf36b8df524e86e29324e679a5d9743450c7714c8ca2165e5  mod=sbi.producer
INFO[0146] Ue [SUPI=imsi-208930000000003] has subscribed Snssai = 1-010203  mod=sbi.producer
INFO[0146] NSSF finds an AMF[amfid=112233] to handle the UE[subscribed slice=1-010203]  mod=sbi.producer
INFO[0147] UeContext 0 is removed from pool              mod=context
INFO[0147] UeContext 0 is terminated                     mod=sbi.producer
WARN[0149] UeContext reaches its end-of-life             mod=sbi.producer
```

- amf:
```bash
INFO[0129] Receive an InitUeContextRequest from /ran/208-93/daejeon_4_0  mod="amf:sbi.producer"
WARN[0129] No UeContext is found for SUCI[suci-0-208-93-0000-0-0-0000000003], create a new UeContext  mod="amf:gmm"
INFO[0129] service /ran/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0130] add a service instance [id=4] to group /ran/208-93/daejeon  mod=registry
INFO[0130] Start handling of NAS RegistrationRequest     mod="amf:gmm"
INFO[0130] Send a NAS SecurityModeCommand                mod="amf:gmm"
INFO[0130] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0130] Handle Security Mode Complete                 mod="amf:secmode"
INFO[0130] Got an IMEISV for the UE: imeisv-4370816125816151  mod="amf:secmode"
INFO[0130] Security mode establishment completed         mod="amf:gmm"
INFO[0130] Send an AmPolicyAssociation request to PCF    mod="amf:gmm"
INFO[0130] Allocate registration areas is not implemented  mod="amf:gmm"
INFO[0130] Assign ladn information for UE is not implemented  mod="amf:gmm"
INFO[0130] Send a InitialContextSetupRequest to Proxy RAN  mod="amf:gmm"
INFO[0131] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0131] Receive a UplinkNasTransport from Proxy RAN   mod="amf:gmm"
INFO[0131] No SmContext found for PDU session 1          mod="amf:gmm"
INFO[0131] Ue[SUPI=imsi-208930000000003] sends a snssai[{%!s(int32=1) 010203}] for the requested session[1]  mod="amf:context"
INFO[0131] Select Dnn = internet for the SmContext       mod="amf:context"
INFO[0131] Create a SMF consumer /smf/internet/208-93/1-010203  mod="amf:context"
INFO[0131] service /smf/internet/208-93/1-010203 is not in the registry, lets query the controller  mod=registry
INFO[0147] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0163] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0179] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0195] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
INFO[0212] Receive a UplinkNasTransport for Ue[CoreNgapId=0]  mod="amf:sbi.producer"
```

- pran:
```bash
INFO[0091] SCTP Accept from: 127.0.0.1:37801             tag="pran:ngap"
INFO[0091] Create a new NG connection for: 127.0.0.1:37801  tag="pran:ngap"
INFO[0091] Receive a  NG Setup request                   tag="pran:ngap"
INFO[0091] Send an NG-Setup response                     tag="pran:ngap"
INFO[0111] Receive an Initial UE Message from gnB        tag="pran:ngap"
WARN[0111] Ue for RanNgapId=1 not found                  tag=ran
INFO[0111] Add UeContext [CuNgapId=0]                    tag="pran:context"
INFO[0111] A new UeContext [RanNgapID: 1, CuNgapId: 0] is created  tag="pran:ngap"
INFO[0111] service /damf/208-93/daejeon is not in the registry, lets query the controller  mod=registry
INFO[0112] add a service instance [id=2] to group /damf/208-93/daejeon  mod=registry
INFO[0112] Receive a Downlink Nas message from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0112] Send a DownlinkNasTransport to gnB for ue[ranNgapId=1]  tag=ue
INFO[0112] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0112] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0112] Receive a InitUeContextStatus from AMF for UE [cuNgapId=0]  mod="pran:sbi.producer"
INFO[0112] service /amf/208-93/112233 is not in the registry, lets query the controller  mod=registry
INFO[0113] add a service instance [id=3] to group /amf/208-93/112233  mod=registry
INFO[0113] Connected state                               tag=ue
INFO[0113] Receive AmfUeId=0                             tag=ue
INFO[0113] Receive a Downlink Nas message from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0113] Connected state                               tag=ue
INFO[0113] Send a DownlinkNasTransport to gnB for ue[ranNgapId=1]  tag=ue
INFO[0113] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0113] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0114] Connected state                               tag=ue
INFO[0114] Receive a InitCtxSetupReq from AMF for UE [id=0]  mod="pran:sbi.producer"
INFO[0114] Connected state                               tag=ue
INFO[0114] Send initial context setup with [CoreNgapId=0][RanNgapId=1] to gnB  tag=ue
INFO[0114] Receive an Initial Context Setup Response from gnB  tag="pran:ngap"
ERRO[0114] an sbi job is timeouted                       tag=ue
ERRO[0114] Fail to send InitCtxSetupReq: Timeout         mod="pran:sbi.producer"
INFO[0114] Connected state                               tag=ue
WARN[0114] Orphan InitCtxSetRsp message                  tag=ue
INFO[0114] Receive an Uplink NAS Transport message from gnB  tag="pran:ngap"
INFO[0114] Forward a NasUplinkTransport to AMF for ue[ranNgapId=1], coreNgapId=0]  tag="pran:ngap"
INFO[0114] Connected state                               tag=ue
ERRO[0124] SCTPRead: interrupted system call             tag="pran:ngap"
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

- pcf:
```bash
WARN[0089] Receive a CreateIndividualAMPolicyAssociation from AMF for ue[SUPI=imsi-208930000000003], return a dummy PolicyAssociation  mod=sbi.producer
```

- gnb:
```bash
[2023-07-05 04:38:37.788] [rrc] [debug] UE[1] new signal detected
[2023-07-05 04:38:37.789] [rrc] [info] RRC Setup for UE[1]
[2023-07-05 04:38:37.789] [ngap] [debug] Initial NAS message received from UE[1]
[2023-07-05 04:38:40.418] [ngap] [debug] Initial Context Setup Request received
[2023-07-05 04:54:31.461] [sctp] [debug] SCTP association shutdown (clientId: 2)
[2023-07-05 04:54:31.461] [sctp] [warning] Unhandled SCTP notification received
[2023-07-05 04:54:31.461] [ngap] [error] Association terminated for AMF[2]
[2023-07-05 04:54:31.461] [ngap] [debug] Removing AMF context[2]
```

- ue:
```bash
UERANSIM v3.2.6
[2023-07-05 04:38:37.777] [nas] [info] UE switches to state [MM-DEREGISTERED/PLMN-SEARCH]
[2023-07-05 04:38:37.788] [rrc] [debug] New signal detected for cell[1], total [1] cells in coverage
[2023-07-05 04:38:37.788] [nas] [info] Selected plmn[208/93]
[2023-07-05 04:38:37.788] [rrc] [info] Selected cell plmn[208/93] tac[1] category[SUITABLE]
[2023-07-05 04:38:37.788] [nas] [info] UE switches to state [MM-DEREGISTERED/PS]
[2023-07-05 04:38:37.788] [nas] [info] UE switches to state [MM-DEREGISTERED/NORMAL-SERVICE]
[2023-07-05 04:38:37.788] [nas] [debug] Initial registration required due to [MM-DEREG-NORMAL-SERVICE]
[2023-07-05 04:38:37.788] [nas] [debug] UAC access attempt is allowed for identity[0], category[MO_sig]
[2023-07-05 04:38:37.788] [nas] [debug] Sending Initial Registration
[2023-07-05 04:38:37.789] [nas] [info] UE switches to state [MM-REGISTER-INITIATED]
[2023-07-05 04:38:37.789] [rrc] [debug] Sending RRC Setup Request
[2023-07-05 04:38:37.789] [rrc] [info] RRC connection established
[2023-07-05 04:38:37.789] [rrc] [info] UE switches to state [RRC-CONNECTED]
[2023-07-05 04:38:37.789] [nas] [info] UE switches to state [CM-CONNECTED]
[2023-07-05 04:38:38.569] [nas] [debug] Authentication Request received
[2023-07-05 04:38:40.007] [nas] [debug] Security Mode Command received
[2023-07-05 04:38:40.007] [nas] [debug] Performing kAMF derivation from kAMF in mobility
[2023-07-05 04:38:40.007] [nas] [debug] Selected integrity[1] ciphering[1]
[2023-07-05 04:38:40.418] [nas] [debug] Registration accept received
[2023-07-05 04:38:40.419] [nas] [info] UE switches to state [MM-REGISTERED/NORMAL-SERVICE]
[2023-07-05 04:38:40.419] [nas] [info] Initial Registration is successful
[2023-07-05 04:38:40.419] [nas] [debug] Sending PDU Session Establishment Request
[2023-07-05 04:38:40.419] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 04:38:56.876] [nas] [warning] Retransmitting PDU Session Establishment Request due to T3580 expiry
[2023-07-05 04:38:56.877] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 04:39:12.989] [nas] [warning] Retransmitting PDU Session Establishment Request due to T3580 expiry
[2023-07-05 04:39:12.989] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 04:39:29.085] [nas] [warning] Retransmitting PDU Session Establishment Request due to T3580 expiry
[2023-07-05 04:39:29.085] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 04:39:45.165] [nas] [warning] Retransmitting PDU Session Establishment Request due to T3580 expiry
[2023-07-05 04:39:45.165] [nas] [debug] UAC access attempt is allowed for identity[1, 2], category[MO_sig]
[2023-07-05 04:40:01.233] [nas] [error] PDU Session Establishment procedure failure, no response from the network after 5 attempts
[2023-07-05 04:40:01.233] [nas] [debug] Aborting SM procedure for PTI[1], PSI[1]
```