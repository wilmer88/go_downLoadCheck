import { Component } from '@angular/core';

@Component({
  selector: 'app-users',
  templateUrl: "./users_list.component.html",
  styles: [
  ]
})
export class usersTable {
  pageTitle: string ="Family Members";
  imageWidth: number = 80;
  imageMargin: number = 5;
  showImage = false;
  listFilter = "doris";
 members= [
    {id: 0, name:"wilmer", imageUrl: "assets/images/wilmer.jpg"},
    {id: 1, name:"doris"},
    {id: 10, name:"felix" },
];
fam: string = "member";
toggleImage():void {
  this.showImage = !this.showImage;
}

  constructor(){}
  // ngOnInit(){
  //   //put code for start up
  // }

}