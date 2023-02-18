import { Component } from '@angular/core';
import { ApiService } from '../api.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
    public message: string = ""
    constructor(private api: ApiService) {};  
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

    /*onSubmit(){
      this.api.getSubs().subscribe( (res: Object) => {
        const response: string = JSON.stringify(res);
        const responseMessage = JSON.parse(response);
          
        let index: number = 0;
        while (responseMessage[index] != null){
          console.log(responseMessage[index]); //get Subcription
          console.log(responseMessage[index]["name"]); //get value from Subcription (see userData for all parameters)
  
          this.allSubs += responseMessage[index]["name"] + " ";
  
          index += 1;
        }
      });
    }*/

    
}
