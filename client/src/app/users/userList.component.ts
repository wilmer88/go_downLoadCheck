import { Component } from '@angular/core';
import {OnInit} from "@angular/core";
import { Input } from '@angular/core';
import { Imembers } from './Iuser';
import { UserService } from './user.service';

@Component({
  selector: 'app-users',
  templateUrl: "./users_list.component.html",
  styleUrls: ["./user_list.component.css"]
})
export class usersTable implements OnInit {
 
  pageTitle: string ="Family Members";
  imageWidth: number = 80;
  imageMargin: number = 5;
  showImage = false;
  errorMassage: string = "";

 private _listFilter: string = "";

 get listFilter(): string {
    return this._listFilter;
 };

 set listFilter(value: string) {
    this._listFilter = value;
    console.log("in setter: ", value);
    this.filterMembers = this.performFilter(value)
 };

 filterMembers: Imembers[]=[];

 members:  Imembers[]= [];

// fam: string = "member";

toggleImage():void {
  this.showImage = !this.showImage;
}
  constructor(private userService:UserService){}

  ngOnInit(): void  {
    this.userService.getUsers().subscribe({
      next: members=>{
        this.members =  members;
        this.filterMembers = this.members;
      },
      error: err => this.errorMassage = err
    })
  };

  onRatingClicked(message:string):void{
    this.pageTitle="product list: " + message;
  };

  performFilter(filterBy: string): Imembers[]{
    filterBy = filterBy.toLowerCase();
    return this.members.filter((member: Imembers) => 
    member.FirstName.toLocaleLowerCase().includes(filterBy)
    )
  }

}