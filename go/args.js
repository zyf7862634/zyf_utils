{
    "SaveDataArgs":{
        "TxData": {
            "curId": "1001",
            "curStep": 1,
            "cryptoData":"cryptodata1111",
            "rootId": "1001",
            "recieveMember":{"member": "member01hash","key": "key01"},
            "notifyMember":[
                {"member":"member02hash", "key": "key02"},
                {"member":"member03hash", "key": "key03"}
            ]
        },
        "PermTable": [
            {
                "stepNo":1,
                "permList":[
                     {"memberhash":"asd", "perm": "r"},
                     {"memberhash":"asdf", "perm": "w"}
                ]
            },
            {
                 "stepNo":2,
                 "permList":[
                        {"memberhash":"asdss", "perm": "rw"},
                         {"memberhash":"asdfw", "perm": "w"}
                 ]
            }
        ]
    }
}
