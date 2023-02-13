// import { HttpClient, HttpErrorResponse } from "@angular/common/http";
// import { Injectable } from "@angular/core";
// import { catchError, Observable, tap, throwError } from "rxjs";
// import { Imembers } from "./user";
// @Injectable({
//     providedIn: "root"
// })
// export class UserService {
//     private userUrl = "http://localhost:3000/users";
//     constructor(private http: HttpClient){}

//     getUsers(): Observable <Imembers[]>{
//         return this.http.get<Imembers[]>(this.userUrl).pipe(
//             tap(data => console.log("ALL",JSON.stringify(data))),
//             catchError(this.handleError)
//         );
//     }
//     private handleError(err: HttpErrorResponse){
//         let errorMassage = "";
//         if(err.error instanceof ErrorEvent) {
//             errorMassage = `an error occurred: ${err.error.message}`
//         }
//         console.error(errorMassage);
//         return throwError(()=> errorMassage)
//     }

// }



// [
//     {ID: 0, Name:"wilmer", ImageUrl: "assets/images/wilmer.jpg", Happiness:3.3},
//     {ID: 1, Name:"doris", ImageUrl: "assets/images/wilmer.jpg", Happiness:4.7},
//     {ID: 10, Name:"felix", ImageUrl: "assets/images/wilmer.jpg", Happiness:4.5 },
// ]