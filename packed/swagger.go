package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDQNA0OYkACIgycDMXlienpqUX6UFovqzg/LzSElYGxP1Y/4c7Ug9ki9iLH3/8Ob/lTd3zpYrnvrzwF+Fn9LznPqrT2nem5TF/ytFLxheY1h3a5vnH8sqAj7MOCjogNP56nLK9vS0nbbjH7zQ4d3Y7iqHu/nz2dPd3/nt0LQSeOq+06ngwvGy5IR1/P2N/JZRlS4iDlK/vQS+pYL/e00i1zvIvTvpVES9/dL7fQa6Fo7hEvaf1bKQFHWdnDX86bvlZU53zwwZmCXm6/ozxfnSz7VRd7eFLfrOfq58/Pr5m7d96N5fnXVy/v1km2KPq76XH4tzfH9xzZ/L3sVx+b2BOeSDa2gq28+2W7+48tsOy6I1tV6W1da2L9T+jGp6SNKTN3M3Nd4jB+FdS169e6LutQDvZWh46XDRI6C5xX/pRL0am4rHv3P4fz9s/1gTzOiq3Wy/uMRJ13lQo3St4I08w53Fhv2b2eJ+3PCz/9Ofq/cwu4eSWWaq2M/jT9cfTeHxbS8/M68r9+OvPPZl9Ryuzcf6UR77MO3OabtTp167+tM/bs35lx/v6F3x8jzBbvebf91l32vGlWLyvOMfYyyD6bu7SpNVjlwIGahF77m4tPdlnv/rCYl2vmD1udwwuaTLY/2tob8vh9r7hg/DL+7+nlyx+W28p/bpPxXTU5tP/Rebcwx0vxagtZn6c2H9iplVKhNenCyf21Gy5f7j0v0+Lw2rn1tXOCvf26ezZ3bjh9v/1m6n4mBob//wO82TnK/kre+cXIwBDJycAAS0QMDFpoiYgdkYjA6WZCrH4CSDeymgBvRiYRZkQiRDYZlAhh4H8jiMSbJBFGYXcKBAgw/Hd0ZWLA4jBWNpA8EwMTQycDA0MzE4gHCAAA//+E5/34IgMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
