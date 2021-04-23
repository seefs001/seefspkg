# seefslib 个人自用工具包

- sync/errgroup 提供带recover和并行数的errgroup，err中包含详细堆栈信息
- xconvertor
    - Bytes2String Bytes2Int BytesToInt64
    - IntToString Int2Bytes Int64ToString Int64ToBytes IntToInt64 ToInt64
    - JSONStrToMap StringToInt64 StringToInt
    - Float64ToString
    - StructToMap StructToJSONStr
- xdate GetCurrntTimeStr GetCurrntTime
- xfile 
    - GetFileSize GetFileExt CheckFileExist 
    - CheckFilePermission IsNotExistMkDir MkDir OpenFile
    - IsBinary IsImg IsDir CopyFile CopyDir
- xrandom RandStringRunes GenRandomCode RandInt
- xstring Len(utf8) Substr JoinInts SplitInts
- xtcp CreateTCPListener CreateTCPConn
- xnet xnet.New()
- xzip Zip UnZip