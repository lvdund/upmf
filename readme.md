## <u> User Plane Managerment Function </u>

1. UPF registration and SMF Query Path 
2. Run:
    ```
    go run cmd/main.go
    ```
3. Network reference point
- N41 - reference point between the SMF and the UPF: This interface is intended for session establishment.
- N42 - reference point between the UPMF and the UPF: This interface is dedicated to communication between the UPF and the UPMF.
- N43 - reference point between the SMF and the UPMF: This interface is used to find the optimal UPFs routing for the SMF.

[Link paper](https://ieeexplore.ieee.org/abstract/document/10392651)
