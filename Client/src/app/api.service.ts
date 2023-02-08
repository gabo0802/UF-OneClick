import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }


  createUser(userData: {email: string, password: string, ussername: string}){
    
    this.http.post('http://localhost:8000/', JSON.stringify(userData)).subscribe((res)=>{
      console.log(res);
    });
  }
}
