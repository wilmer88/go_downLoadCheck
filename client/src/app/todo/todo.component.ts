import { Component } from '@angular/core';
import { UserService } from './user.service';

@Component({
  selector: 'app-todo',
  template: `
    <p>
      todo works!
    </p>
    <div>
  <ul>
  <li *ngFor="let person of persons" (click)="seePerson(person)">{{person.id}}</li>
  </ul>
   </div>

   <div *ngIf="selectedPerson">
   youve selected {{this.selectedPerson.name}}
   </div>


  `,
  styles: [
  ]
})
export class TodoComponent {
  persons = [
    {id: 0, name:"wilmer"},
    {id: 1, name:"doris"},
    {id: 10, name:"felix"},
];

constructor(private userService: UserService){}

// ngOnInit(){ this.persons = this.userService.getUsers()}

title: string = "person"
  selectedPerson: any;

seePerson(person: any ){
  this.selectedPerson = person
  alert("clicked person: " + this.selectedPerson.name)
  console.log(this.selectedPerson.name)
}

}
