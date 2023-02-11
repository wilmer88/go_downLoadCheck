import { Component } from '@angular/core';

@Component({
  selector: 'CoatnCode-root',
  template: `
   
    <div style="text-align:center" class="content">
      <h1>
        Welcome to {{title}}!
      </h1>
     <app-jobs></app-jobs>
     <app-todo></app-todo>

    </div>
    <router-outlet></router-outlet>
  `,
  styles: []
})
export class AppComponent {
  title = 'client';
}
