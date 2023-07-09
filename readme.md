## <u> User Plane Managerment Function </u>

1. UPF registration and SMF Query Path 
2. Run:
    ```
    go run cmd/main.go
    ```
3. For sbi test, join this link [**postman**](https://app.getpostman.com/join-team?invite_code=0e108cef4305edc1654d849cb7ae4196&target_code=2ab452511741d2b0e2d60797b64c4e61)

4. Network reference point
- N41 - reference point between the SMF and the UPF: This interface is intended for session establishment.
- N42 - reference point between the UPMF and the UPF: This interface is dedicated to communication between the UPF and the UPMF.
- N43 - reference point between the SMF and the UPMF: This interface is used to find the optimal UPFs routing for the SMF.