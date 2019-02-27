package main

import "fmt"

type Pages struct {
	Pages []int
	Current int
}

func (p *Pages) FirstPage() int{
	p.Current=p.Pages[0]
	return p.Current
}

func (p *Pages) LastPage() int{
	p.Current=p.Pages[len(p.Pages)-1]
	return p.Current
}

func (p *Pages) GetCurrentPageNumber() int{
	return p.Current
}

func (p Pages) HasNext() bool{
	if p.Current==p.Pages[len(p.Pages)-1]{
		return false
	}
	return true
}

func (p Pages) HasPrevious() bool{
	if p.Current==p.Pages[0]{
		return false
	}
	return true
}

//func (p *Pages) Next() int{
//	return p.Current
//}
//
//func (p *Pages) Previous() int{
//	return p.Current
//}

func main()  {
	s:=Pages{[]int {1,2,3,4,5,6},4}
	fmt.Println(s)
	fmt.Println(s.FirstPage())
	fmt.Println(s)
	fmt.Println(s.LastPage())
	fmt.Println(s)
	fmt.Println(s.GetCurrentPageNumber())
	fmt.Println(s.HasNext())
	fmt.Println(s.HasPrevious())
}