package stack

type Stack struct {
	i int
	data [10]int
}

func (s *Stack) Xpush(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Xpop() (ret int) {
	ret = s.data[0]
	if (s.i>0) {
		s.i--
		for k:=1;k<len(s.data)-1;k++ {
			if (s.data[k-1] == 0) { //@todo go初始化时就是0值，这个还不知道如何解
				break
			}
			s.data[k-1] = s.data[k]
		}
	}
	return ret
}

func (s *Stack) Xget() (int,[10]int) {
	return s.i,s.data
}
