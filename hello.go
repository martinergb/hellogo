// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package main

import (
	"fmt"
	"log"
	"time"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func main() {
	for {
		oldBook := Person{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*Person_PhoneNumber{
				{Number: "555-4321", Type: Person_HOME},
			},
		}

		fmt.Println(oldBook)
		out, err := proto.Marshal(&oldBook)

		if err != nil {
			log.Fatalln("Failed to encode address book:", err)
		} else {
			fmt.Println(out)
		}

		//
		newBook := &Person{}

		err = proto.Unmarshal(out, newBook)

		if err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}

		fmt.Println(*newBook)

		time.Sleep(time.Millisecond * 5)
	}
}
