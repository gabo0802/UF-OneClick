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
    public adminButtonVisible: number = 0;
    constructor(private api: ApiService, private router: Router, private authService: AuthService) {};  
    createSubFrom: FormGroup = {} as FormGroup;
    removeSubFrom: FormGroup = {} as FormGroup;
    newsLetterForm: FormGroup = {} as FormGroup;
    createCustomSubFrom: FormGroup = {} as FormGroup;
    createSubServiceForm: FormGroup = {} as FormGroup;

    ngOnInit(){
      if (document.cookie.includes("currentUserID=1")){
        this.adminButtonVisible = 100
      }else{
        this.adminButtonVisible = 0
      }

      this.createSubServiceForm = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
        'price': new FormControl(null, [Validators.required, Validators.pattern('^[$]{0,1}[0-9]{1,4}.[0-9][0-9]+$')]),
      });

      this.createSubFrom = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
      });

      this.removeSubFrom = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
      });

      this.newsLetterForm = new FormGroup({
        'name': new FormControl(null, null),
      });

      this.createCustomSubFrom = new FormGroup({
        'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),        
        'dateadded': new FormControl(null, [Validators.required, Validators.pattern('^[0-2][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]$')]),
        'dateremoved': new FormControl(null, Validators.pattern('^[0-2][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]$')),
      });

      this.getUserSubscriptions()
      //this.userID = int(document.cookie);
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
          allSubsString += "All Active Subscriptions <br>"

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

    onSubmit4(){
      this.api.createSub(this.createSubServiceForm.value).subscribe( (resultMessage: string[]) => {
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

    onSubmit3(){
      //Prime Video
      //2023-01-02 15:04:05
      this.api.addOldUserSub(this.createCustomSubFrom.value).subscribe( (resultMessage: string[]) => {
        location.reload();
        alert(resultMessage[0] + ": " + resultMessage[1])
      })
    }

    public doLogout(){
      this.api.logout().subscribe( (res) => {
        alert("Logging Out...")
        this.authService.userLogOut();
        this.router.navigate(['login']);
      })
    }

    public getUserData(){
      this.api.getAllUserData().subscribe( (res) => {
        var allSubsString:string = ""
        const response: string = JSON.stringify(res);
        const responseMessage = JSON.parse(response);

        console.log(responseMessage)
        console.log(responseMessage[1])

        if (responseMessage["Error"] == undefined){
          allSubsString += "All User Data:<br>";
          var previousName = "";

          let index: number = 1;
          while (responseMessage[index] != null){
            if(responseMessage[index]["dateadded"] != ""){
              if (previousName != "User Subscriptions"){
                allSubsString += "<br> All User Subscriptions: <br>";
                previousName = "User Subscriptions"
              }

              if (responseMessage[index]["dateremoved"] == ""){
                var dateAdded: string = responseMessage[index]["dateadded"]
                allSubsString += "[" + responseMessage[index]["email"] + " "+responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length) + "] <br>";
              }else{
                var dateAdded: string = responseMessage[index]["dateadded"]
                var dateRemoved: string = responseMessage[index]["dateremoved"]
                allSubsString += "[" + responseMessage[index]["email"] + " "+ responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length) + " to " + dateRemoved.substring(0, dateRemoved.length) + "] <br>";
              }
            }else if (responseMessage[index]["price"] != ""){
              if (previousName != "Subscriptions"){
                allSubsString += "<br> All Subscriptions: <br>";
                previousName = "Subscriptions"
              }

              allSubsString += "[" + responseMessage[index]["subid"] + " " + responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + "] <br>";
            }else{
              if (previousName != "Users"){
                allSubsString += "<br> All Users: <br>";
                previousName = "Users"
              }
              console.log("User", index)
              console.log(responseMessage[index])

              allSubsString += "[" + responseMessage[index]["userid"] + " "+ responseMessage[index]["email"] + "] <br>";
            }
            
            index += 2
          }

          console.log(allSubsString)
          document.getElementById("allSubs")!.innerHTML = allSubsString;
        }
      })
    }

    public resetAll(){
      alert("Reset Starting...")

      this.api.resetall().subscribe( (res) => {
        alert("Reset Successful! Logging Out...")
        this.authService.userLogOut();
        this.router.navigate(['login']);
      })
    }
    
    newsLetter(){
        var message:string = (document.getElementById('newslettermessage') as HTMLInputElement).value;

        if(message != "Enter Message For Newsletter"){
          this.newsLetterForm.controls['name'].setValue(message);
          (document.getElementById('newslettermessage') as HTMLInputElement).value = "Enter Message For Newsletter";
          
          this.api.sendNews(this.newsLetterForm.value).subscribe( (res) => {
              alert("Message Sent")
          })
        }
    }
}
