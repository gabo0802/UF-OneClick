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
    createSubFrom: FormGroup = {} as FormGroup;
    removeSubFrom: FormGroup = {} as FormGroup;

    ngOnInit(){
      this.createSubFrom = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
      });

      this.removeSubFrom = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
      });

      this.getUserSubscriptions()
    }

    /*getDefaultSubscriptions(){
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
    }*/

    getUserSubscriptions(){
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

    onSubmit(){
      this.api.addUserSub(this.createSubFrom.value).subscribe( (resultMessage: string[]) => {
        location.reload();
        alert(resultMessage[0] + ": " + resultMessage[1])
      })
    }

    onSubmit2(){
      this.api.removeUserSub(this.removeSubFrom.value).subscribe( (resultMessage: string[]) => {
        location.reload();
        alert(resultMessage[0] + ": " + resultMessage[1])
      })
    }

    public doLogout(){
      this.api.logout().subscribe( (res) => {
        this.authService.userLogOut();
        this.router.navigate(['login']);
      })
    }
    
}
