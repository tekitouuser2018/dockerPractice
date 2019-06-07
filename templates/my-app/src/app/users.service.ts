import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../model';

const httpOptions = {
  headers  : new HttpHeaders({
    // 'Content-Type': 'application/x-www-form-urlencoded',
    // 'Content-Type': 'application/json',
    // 'Access-Control-Allow-Origin': '*',
    // 'Access-Control-Allow-Headers': 'Origin',
    // add_header Access-Control-Allow-Origin '*' always;
    // 'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Authorization, Accept, Content-Type, User-Agent, Keep-Alive',
    // 'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    // 'Access-Control-Allow-Credentials': 'true'
  })
}

@Injectable({
  providedIn:"root"
})
export class UsersService {
  headers :HttpHeaders;

  ngOnInit() {
    // POST で使うヘッダーを作っておく
    // this.headers = new HttpHeaders();
    // this.headers.append('Content-Type', 'application/json');
    // this.headers.append('Content-Type', 'application/x-www-form-urlencoded');
  }

  constructor(private http: HttpClient) { }

  getUsers(): Observable<User[]>{
    console.log('-----getUsers method called-------');
    // return this.http.get<User[]>(`http://docker.for.mac.localhost:8080/`,httpOptions);
    return this.http.get<User[]>(`http://docker.for.mac.localhost:6969/test2/`);

  }
}
