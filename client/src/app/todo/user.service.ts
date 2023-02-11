import {Injectable} from '@angular/core';

@Injectable({ providedIn: 'root' })

export class UserService {
    // constructor(private http: HttpClient){}
    getUsers(){
        // return this.http.get<[{id:0, name:"wilmer"}]>
        // return this.http.get<User[]>('/users');
    }
}