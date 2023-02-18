import { Component } from '@angular/core';
import { ApiService } from '../api.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
    public message: string = ""
    constructor(private api: ApiService, private router: Router, private authService: AuthService) {};  
    loginForm: FormGroup = {} as FormGroup;

    ngOnInit(){
      this.api.getSubs().subscribe( (res: Object) => {
        var allSubsString:string = ""
        const response: string = JSON.stringify(res);
        const responseMessage = JSON.parse(response);
        
        if (responseMessage["Error"] == undefined){
          let index: number = 0;
          while (responseMessage[index] != null){
            if (responseMessage[index]["dateremoved"] == ""){
              var dateAdded: string = responseMessage[index]["dateadded"]
              allSubsString += "[" + responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length - 9) + "] <br>";
            }

            index += 1;
          }

          document.getElementById("allSubs")!.innerHTML = allSubsString;
        }
      });
    }

    public doLogout(){
      this.api.logout().subscribe( (res: Object) => {
        const response: string = JSON.stringify(res);
        const responseMessage = JSON.parse(response);
        
        if (responseMessage["Error"] == undefined){
          this.authService.userLogOut();
          this.router.navigate(['login']);
        }
      });
    }
    
}
