package examplestructs

import "fmt"

//student => name, ,id, roll_no, address, phone

// type address struct {

// 	address_type string
// 	street string
// 	houseno string
// 	location string

// }

type student struct {
     id   int
	name   string
	roll_no  int
	address   map[string] string
	phone  [] string
	
}

func creatStudent(id int , name string) *student {
             s := student {

				id : id,
				name: name,
			 }
        return &s
}
    func  DisplayStudent(s student) {

		fmt.Println("id =", s.id)
		fmt.Println("name =", s.name)
		fmt.Println("roll_no =", s.roll_no)
		fmt.Println("address =", s.address)
	
	}


	

func InitExampleStructs() {

	std1 := student {

		id : 1,
		name : "ram",
		roll_no: 1,
		address: map[string]string{"home": "buddhanagar", "Office": "pulchowk"},
		phone: []string {"2959595","2996269"},
	}
	fmt.Printf("%+v\n", std1)
	std2 := student{

		id : 1,
		name : "ram",
	}
	std2.roll_no = 2
	std2.address = map[string]string{"home": "x", "y": "y"}
    fmt.Printf("%+v\n", std2)
	fmt.Println(std2.name)
	fmt.Println(std2.address["home"])

   std3 := creatStudent(3, "shyam")
   fmt.Println(std3)
   DisplayStudent(std1)

//    var a = 10
//    var p *int
//    p = &a
//    fmt.Println(*p)

}