package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAINg4OYkACIgycDMXlienpqUX6UFovqzg/LzSElYFRPko/YU2IYf9lA5F7v7V1DkeYHrjBa7hR4YpoacHj61dCk5Z/417hFDmDKyiz9QgLt9MK5tMv26Tf1fnf+tczefKWXVfy1h1O/XPv+zz7+vfFduUcii2z/HW6HC7bLVxfpZ+x3alVRvuWuuhyFyfVia0ZB/e4fJ41y2jODDXzHz+MLmi1sUVNnOrIvlYr6qMOi/XL5mBdSR3fZEEjES+24MUqpRmH11mLR14xKt7/c+P/qr1/9+rs+7szZ97JSS0vd35dHfFH613ea8H7/0VX779yeYLRWSGHC89Td5TJ2J062BK15Nul5e5iq/Vvzw+fp9x/dpJqGbuICtec7YaulrVLKrS9D+sIS1m8Fb4iucXA16KxpdL87gPb6wbb1O25LXpaRb7NU/MQad/x3Y3ZWWKFBJtB8fTiH0X5R13bZev3znzGbjRr5rfMoke9n7eYxPQuE3jUOdXskWaGSMIf/4ajTT+9jwi3Xp/m+0pD55/466czr8qGSl5xcls3eaf9VoWPuvsjzv3e1fhGzHyqxsH5QRP+zt+z+lPNzAWOdws2ZOy9k3tu77Tjolvbl0++z87A8P9/gDc7h/7CNuHZjAwMfqwMDLBYZmAIRYtldkQsgyNWKUo/AaQbWU2ANyOTCDMilSCbDEolMPC/EUTiTTMIo7A7BQIEGP47PmNkwOIwVjaQPBMDE0MnAwODChOIBwgAAP//soCxWMMCAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
