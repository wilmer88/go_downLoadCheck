import { Component } from '@angular/core';

@Component({
  selector: 'CoatnCode-root',
  template: `
  <div class="container text-center">
  <div class="row align-items-start">

    <div class="col-2">
    <app-jobs></app-jobs>
    </div>

    <div class="col">
    <div style="text-align:center" class="content">
    
    <h1>Welcome to {{title}}!</h1>
    <app-users></app-users>
    </div>
    </div>

    <div class="col-2">
    <app-todo></app-todo>
    </div>


  </div>
</div>
   

    <router-outlet></router-outlet>
  `,
  styles: []
})
export class AppComponent {
  title = 'client';
}
