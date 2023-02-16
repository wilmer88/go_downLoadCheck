# Client

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.1.5.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.

- <img [src]="customer.imagePath" />
- <button [disabled]="!isEnabld">save</button>
- <div [style.color]="textColor" [attr.aria-label]="text">...</div>

- <div class="button" [ngClass]="{foo:isActive}, bar: isDisabled" [attr.aria-label]="text">...</div>

- <div class="field">
<label class="label" for="firstName">first name</label>
<input class="input" id="firstName" [(ngModel)]="jobs.firstName"/>
 <button (click)="onJobs">save</button>
 <div *ngIf="selectedJob">
 You selected {{selectedJob.title}}
  <ul *ngFor="let job of jobs">
 <li>Title:{{job.title}} Location:{{job.location}} </li>
 </ul>
</div>
 </div>

 -  person <pre>{{persons | json}}</pre>

 {{model}} interpolation
`[property]` Bind to a DOM proberty
(event) bind to an event
` ngModel ` 2 way data binding
` *ngIf ` conditional element
` *ngFor ` loop


<input type="text" value="firstName" [(ngModel)]="title" />

- ` go mod tidy `
 - ` ng generate component jobs `
 - ` ng serve `
 - ` ng new client --style scss --routing -s -t ` 
 - https://angular.io/errors


- ` go mod tidy `
 - ` ng generate component jobs `
 - ` ng serve `
 - ` ng new client --style scss --routing -s -t `
 ` git remote -v `
 `heroku repo:purge_cache -a appname`
 `git push <heroku url> HEAD:master`
 `heroku buildpacks:set heroku/php`
 `go mod vendor`