import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit{

  constructor(private api: ApiService, private dialogs: DialogsService) {}

  username: string = '';
  userSubscriptions: Subscription[] = [];

  ngOnInit(): void {

    this.api.getUserInfo().subscribe({

      next: (res: Object) => {

        let data = JSON.stringify(res);
        let userData = JSON.parse(data);

        this.username = userData.username;
      },
      error: (error: HttpErrorResponse) => {

        this.dialogs.errorDialog("Unexpected Error!", "Please try again later.");
      }
    });

    this.api.getUserSubscriptions().subscribe({

      next: (res: Subscription[]) => {
        this.userSubscriptions = res;
      },
      error: (error: HttpErrorResponse) => {

        this.dialogs.errorDialog("Unexpected Error!", "Please try again later.");
      }
    });
  }
}



































































// //Multi-Line Comments Exist Just Do /* */

//   // public message: string = ""
//   //   adminButtonVisible: boolean = true;

//     public currentUsername: string = "Loading..."
//     constructor(private api: ApiService, private router: Router, private authService: AuthService) {};  
//     public currentMode = 1;
//     createUserSubForm: FormGroup = {} as FormGroup;
//     removeUserSubForm: FormGroup = {} as FormGroup;
//     //newsLetterForm: FormGroup = {} as FormGroup;
//     createCustomUserSubForm: FormGroup = {} as FormGroup;
//     createSubServiceForm: FormGroup = {} as FormGroup;

//      ngOnInit(){
//   //    if (document.cookie.includes("currentUserID=1")){
//   //       this.adminButtonVisible = false
//   //     }else{
//   //       this.adminButtonVisible = true
//   //     }

//       // this.api.getEmailandUsername().subscribe((resultMessage: string[]) => {
//       //   this.currentUsername  = resultMessage[0]; 
//       // });

//        this.createSubServiceForm = new FormGroup({
//          'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
//          'price': new FormControl(null, [Validators.required, Validators.pattern('^[$]{0,1}[0-9]+.99$')]),
//        });

//        this.createUserSubForm = new FormGroup({
//          'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
//        });

//        this.removeUserSubForm = new FormGroup({
//          'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),
//        });

//   //     this.newsLetterForm = new FormGroup({
//   //       'name': new FormControl(null, null),
//   //     });

//        this.createCustomUserSubForm = new FormGroup({
//          'name': new FormControl(null, [Validators.required, Validators.pattern('^[A-z+() ]+$')]),        
//          'dateadded': new FormControl(null, [Validators.required, Validators.pattern('^[0-2][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]$')]),
//          'dateremoved': new FormControl(null, Validators.pattern('^[0-2][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]$')),
//        });

//        this.getUserSubscriptions()
//        this.getLongestUserSubscription()
//   //     //this.userID = int(document.cookie);
//      }

//     /*getDefaultSubscriptions(){
//       this.api.getSubs().subscribe( (res: Object) => {
//         var allSubsString:string = ""
//         const response: string = JSON.stringify(res);
//         const responseMessage = JSON.parse(response);
        
//         if (responseMessage["Error"] == undefined){
//           let index: number = 0;
//           while (responseMessage[index] != null){
//             if (responseMessage[index]["dateremoved"] == ""){
//               var dateAdded: string = responseMessage[index]["dateadded"]
//               allSubsString += "[" + responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length - 9) + "] <br>";
//             }

//             index += 1;
//           }

//           document.getElementById("allSubs")!.innerHTML = allSubsString;
//         }
//       });
//     }*/

//      getUserSubscriptions(){
//         this.api.post_request__with_data({username: "", email: "", password: "", name: "", price: "", dateadded:"", dateremoved:""}, "/api/subscriptions").subscribe( (res: Object) => {
//           var allSubsString:string = ""
//           const response: string = JSON.stringify(res);
//           const responseMessage = JSON.parse(response);
        
//          if (responseMessage["Error"] == undefined){
//            allSubsString += "All Active Subscriptions <br>"

//            let index: number = 0;
//            while (responseMessage[index] != null){
//              if (responseMessage[index]["dateremoved"] == ""){
//                var dateAdded: string = responseMessage[index]["dateadded"]
//                allSubsString += "[" + responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length) + "] <br>";
//            }

//              index += 1;
//            }

//            document.getElementById("allSubs")!.innerHTML = allSubsString;
//          }
//        });
//      }

//      getLongestUserSubscription(){
//       var currentURL: string = "/api/longestsub";
//       var allSubsString:string = "<br>"

//       switch(this.currentMode) {
//         case 1:
//           currentURL = "/api/longestsub";
//           allSubsString += "Longest Subscription: "
//           break;
//         case 2:
//           currentURL = "/api/longestcontinoussub";
//           allSubsString += "Longest Continuous Subscription: "
//           break;
//         case 3:
//           currentURL = "/api/longestactivesub";
//           allSubsString += "Longest Currently Active Subscription: "
//           break;
//         default:
//           currentURL = "/api/longestsub";
//           allSubsString += "Longest Subscription: "
//       }

//       this.api.post_request__with_data({username: "", email: "", password: "", name: "", price: "", dateadded:"", dateremoved:""}, currentURL).subscribe( (res: Object) => {
//         const response: string = JSON.stringify(res);
//         const responseMessage = JSON.parse(response);
      
//         if (responseMessage["Error"] == undefined){
//             allSubsString += responseMessage[0] + "<br>" + responseMessage[1];
//          }

//          document.getElementById("longestSub")!.innerHTML = allSubsString;
//      });
//    }

//    public left(){
//     if (this.currentMode == 1){
//       this.currentMode = 3;
//     }else{
//       this.currentMode--;
//     }
//     //console.log(this.currentMode)
//     this.getLongestUserSubscription()
//   }

//    public right(){
//     if (this.currentMode == 3){
//       this.currentMode = 1;
//     }else{
//       this.currentMode++;
//     }
//     //console.log(this.currentMode)
//     this.getLongestUserSubscription()
//   }

//      newUserSubscription(){
//        this.api.post_request__with_data(this.createUserSubForm.value, "/api/subscriptions/addsubscription").subscribe( (resultMessage: string[]) => {
//          alert(resultMessage[0] + ": " + resultMessage[1])
//          location.reload();
//          //alert(resultMessage[0] + ": " + resultMessage[1])
//        })
//      }

//      newSubscription(){
//        this.createSubServiceForm.controls['price'].setValue( this.createSubServiceForm.controls['price'].value.replaceAll("$", ""));

//        this.api.post_request__with_data(this.createSubServiceForm.value, "/api/subscriptions/createsubscription").subscribe( (resultMessage: string[]) => {
//          alert(resultMessage[0] + ": " + resultMessage[1])
//          location.reload();
//          //alert(resultMessage[0] + ": " + resultMessage[1])
//        })
//      }

//      removeUserSubscription(){
//        this.api.post_request__with_data(this.removeUserSubForm.value, "/api/subscriptions/cancelsubscription").subscribe( (resultMessage: string[]) => {
//          alert(resultMessage[0] + ": " + resultMessage[1])
//          location.reload();
//          //alert(resultMessage[0] + ": " + resultMessage[1])
//        })
//      }

//      customUserSubscription(){
//        this.api.post_request__with_data(this.createCustomUserSubForm.value, "/api/subscriptions/addoldsubscription").subscribe( (resultMessage: string[]) => {
//          alert(resultMessage[0] + ": " + resultMessage[1])
//          location.reload();
//          //alert(resultMessage[0] + ": " + resultMessage[1])
//        })
//      }

    // public doLogout(){
    //   this.api.post_request({name: ""}, "/api/logout").subscribe((res) => {
    //     alert("Logging Out...")
    //     this.authService.userLogOut();
    //     this.router.navigate(['login']);
    //   })
    // }

    // public getUserData(){
    //   this.api.post_request__with_data({username: "", email: "", password: "", name: "", price: "", dateadded:"", dateremoved:""}, "/api/alldata").subscribe( (res) => {
    //     var allSubsString:string = ""
    //     const response: string = JSON.stringify(res);
    //     const responseMessage = JSON.parse(response);

    //     console.log(responseMessage)
    //     console.log(responseMessage[1])

    //     if (responseMessage["Error"] == undefined){
    //       allSubsString += "All User Data:<br>";
    //       var previousName = "";

    //       let index: number = 1;
    //       while (responseMessage[index] != null){
    //         if(responseMessage[index]["dateadded"] != ""){
    //           if (previousName != "User Subscriptions"){
    //             allSubsString += "<br> All User Subscriptions: <br>";
    //             previousName = "User Subscriptions"
    //           }

    //           if (responseMessage[index]["dateremoved"] == ""){
    //             var dateAdded: string = responseMessage[index]["dateadded"]
    //             allSubsString += "[" + responseMessage[index]["email"] + " "+responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length) + "] <br>";
    //           }else{
    //             var dateAdded: string = responseMessage[index]["dateadded"]
    //             var dateRemoved: string = responseMessage[index]["dateremoved"]
    //             allSubsString += "[" + responseMessage[index]["email"] + " "+ responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + " " + dateAdded.substring(0, dateAdded.length) + " to " + dateRemoved.substring(0, dateRemoved.length) + "] <br>";
    //           }
    //         }else if (responseMessage[index]["price"] != ""){
    //           if (previousName != "Subscriptions"){
    //             allSubsString += "<br> All Subscriptions: <br>";
    //             previousName = "Subscriptions"
    //           }

    //           allSubsString += "[" + responseMessage[index]["subid"] + " " + responseMessage[index]["name"] + " $"+ responseMessage[index]["price"] + "] <br>";
    //         }else{
    //           if (previousName != "Users"){
    //             allSubsString += "<br> All Users: <br>";
    //             previousName = "Users"
    //           }
    //           console.log("User", index)
    //           console.log(responseMessage[index])

    //           allSubsString += "[" + responseMessage[index]["userid"] + " "+ responseMessage[index]["username"] + " "+ responseMessage[index]["email"] + "] <br>";
    //         }
            
    //         index += 2
    //       }

    //       console.log(allSubsString)
    //       document.getElementById("allSubs")!.innerHTML = allSubsString;
    //     }
    //   })
    // }

    // public resetAll(){
    //   alert("Reset Starting...")

    //   this.api.post_request({name: ""}, "/api/reset").subscribe( (res) => {
    //     alert("Reset Successful! Logging Out...")
    //     this.authService.userLogOut();
    //     this.router.navigate(['login']);
    //   })
    // }
    
    // newsLetter(){
    //     var message:string = (document.getElementById('newslettermessage') as HTMLInputElement).value;

    //     if(message != "Enter Message For Newsletter"){
    //       this.newsLetterForm.controls['name'].setValue(message);
    //       (document.getElementById('newslettermessage') as HTMLInputElement).value = "Enter Message For Newsletter";
          
    //       this.api.post_request(this.newsLetterForm.value, "/api/news").subscribe( (res) => {
    //           alert("Message Sent")
    //       })
    //     }
    // }
