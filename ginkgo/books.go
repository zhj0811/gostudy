package books

type Book struct{
	Title string 
	Author string 
	Pages int
}

func (book *Book) CategoryByLength() string {
	if (book.Pages>300){
		return "NOVEL"
	}
	return "SHORT STORY"
}
