f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
_, err = f.WriteAt([]byte(content), n)
f.Close()

func main(){	
	f, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND|O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for  i := 0 ; i< 100 ;i++ {		
		msg := fmt.Sprintf("%d\n",i)
		_, err = f.Write([]byte(msg))
		
	}
	f.Close()
}