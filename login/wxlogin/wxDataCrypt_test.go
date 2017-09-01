package wxlogin

import "testing"

func TestNewWXDataCrypt(t *testing.T) {
	wxDataCrypt:=NewWXDataCrypt("111","333")
	if wxDataCrypt ==nil{
		t.Error("TestNewWXDataCrypt error")
	}
	if wxDataCrypt.appId !="111" {
		t.Errorf("appId %s,want 111",wxDataCrypt.appId)
	}
	if wxDataCrypt.sessionKey !="333"{
		t.Errorf("sessionKey %s,want 333",wxDataCrypt.sessionKey)
	}
}

func Test_decryptData(t *testing.T)  {
	wxDataCrypt:=NewWXDataCrypt("111","333")
	if wxDataCrypt ==nil{
		t.Error("TestNewWXDataCrypt error")
	}

	wxed ,err:=wxDataCrypt.decryptData("7cu+ymVDx66Jt6IefBjr1KYa/eYYS/iB/x+UQQ9V7iC4YQSREDyYAEQswYyUnrWWQht1ys+6XZ1HUMWhEe3zJCenFKtUB+mmPcApk5Lpc6tcIqTWKNMnlgaYm9z7v7rYP4YqjvuC2Btd9yiS+7C3Zma5fKgwjw6JKzsVdztT3472VsP82EMhKBgj4D1gALAAcbWlK6EMd/47oqrZQoGbnNSiZ1nqhKiTsGanVN5CGfD+WfQZqItv7wlKofR346dJN91zfOBlePClF/5s13wTPyE/W/2ovrpekLzj/nJlGqVNiSzWCutaGutHqJ/ZfV7VEw92zg1HMaRDLpL1WcpcLs7GCOlfgAmXS33h+ESLBbXZdXlCCKdfMoZUWpwwCT9JnOsXW30Qlzr9Y4WPIHreXrFM8lmg/UpjCWN2vAwW01kfnNFOjbbN5HmpaAK00TvrWYpDfey2Ka4xX2zT7BkvEQ==","z6ifDkyJMJSVA1JFvYHLxg==")
	if err !=nil{
		t.Error("decryptData error")
	}
	t.Log(wxed)
}