dao层

return errors.Wrapf(errors.NotFound, fmt.Sprintf("sql:%s error:%v",sql, err))


业务层

if errors.Is(err,errors.NotFound){

}

if errors.Reason(err, xxxxx) == xxxxx{
	
}