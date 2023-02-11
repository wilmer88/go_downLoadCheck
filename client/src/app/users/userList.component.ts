import { Component } from '@angular/core';
import {OnInit} from "@angular/core";
import { Imembers } from './user';

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

members: Imembers[]= [
    {id: 0, name:"wilmer", imageUrl: "assets/images/wilmer.jpg", happiness:3.3},
    {id: 1, name:"doris", imageUrl: "assets/images/wilmer.jpg", happiness:4.7},
    {id: 10, name:"felix", imageUrl: "assets/images/wilmer.jpg", happiness:4.5 },
];
fam: string = "member";

toggleImage():void {
  this.showImage = !this.showImage;
}

  constructor(){}
  ngOnInit(): void  {
    this.listFilter= ""

    //put code for start up
  }

  performFilter(filterBy: string): Imembers[]{
    filterBy = filterBy.toLowerCase();
    return this.members.filter((member: Imembers) => 
    member.name.toLocaleLowerCase().includes(filterBy)
    )
  }

}