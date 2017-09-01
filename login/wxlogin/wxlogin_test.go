package wxlogin

import "testing"

func TestWXLogin_Login(t *testing.T) {

	wx,err:=New("wx2c6166db1e131d73","0bfdf50511525da3970591dea2ac9627")
	if err !=nil{
		t.Error("WXLogin_Login new error")
	}
	encryptedData:="M2k34oieBEP++DtYe9wjZPILRHhNHCVXpVWKF1DFsAXLlMHZSunTlkJhLQ6xPJcuxSV1XhNuSFUAQopp9SL00bLGPZvy4MCFcJ6aU4UHxcS+BVocktCzeHqnzof5RjiSs4YKCiI59i90TgqnHSTp689B6wSwfFI36e0k/Got9mm2RGp0T7cF0AyiwGDFqj6BWsv2s05b8OkkIy2g+Riweczw5+My0/th8xBzkfHClxkF795ED6M9/eDm9iVuNty6vBYvuwXFa736SO051nVqC/NBSKtzr/sTVJLXL4XTv41jQGeHMP/Ul8JYfPAOsLrrXlk8vmp0tySP8rSg1YXCQhfPtzXB2lBkgJBIyi61Psin3ePeljTRWIF3EauzelrrYVdVBd6V32MENkvag6t+yN0CuHkqiKcyzUW/8lgi2GjIg1tzgRF/Y98QyKD03Sz7VdJE4DWIMy2H2ci2ieNjzw=="
	iv:="aDgf6JXVwdcAHLbEf9eKvw=="
	code:="001CYvVd1ySxbv0RaASd1lKyVd1CYvVH"

	wxu,err := wx.Login(encryptedData,iv,code)
	if err !=nil{
		t.Errorf("WXLogin_Login error %s",err)
		return
	}
	if wxu ==nil{
		t.Error("WXLogin_Login error")
		return
	}
	t.Logf("%+v",wxu)
}
