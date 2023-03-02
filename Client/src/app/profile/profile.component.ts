import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from 'src/app/api.service';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

  constructor(private router: Router, private dialog: MatDialog, private api: ApiService, private authService: AuthService) {}

  hide: boolean = true;
  passwordCharacterLength: number = 3;

  timezoneEdit: boolean = false;

  username: string = '';
  email: string = '';
  timezone: string = '';

  timezoneForm = new FormGroup({
    'timezonedifference': new FormControl(this.timezone, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')]),
  });

  ngOnInit(){
      this.api.getEmailandUsername().subscribe((resultMessage: string[]) => {
        this.username = resultMessage[0], 
        this.email = resultMessage[1];

      });

      this.api.getTimezone().subscribe((resultMessage: string[]) => {
        this.timezone = resultMessage[1] + "UTC";

        this.timezoneForm = new FormGroup({
          'timezonedifference': new FormControl(this.timezone, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')]),
        });
      
      });
  }


  editTimezone(): void{
    this.timezoneEdit = !this.timezoneEdit;

    if (this.timezoneEdit == false){
      var oldTimezone: string = this.timezone
      this.timezone = this.timezoneForm.controls['timezonedifference'].value || this.timezone

      if (this.timezone != oldTimezone ){
        if (confirm("Are you sure you want to change your time from " + oldTimezone + " to " + this.timezone + "?")){
          var actualtimezone: string = this.timezone.replaceAll("UTC", "")
          actualtimezone  = actualtimezone.replaceAll("+", "")

          this.api.updateTimezone(actualtimezone).subscribe((res) => {
            alert("Timezone Changed!")
          })
        }
      }
    }

    this.timezoneForm = new FormGroup({
      'timezonedifference': new FormControl(this.timezone, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')]),
    });
  }  
 
  back(): void {
    this.router.navigate(['users']);
  }

  delete(): void {
    if (confirm("Are you sure you want to delete user " + this.username + "?")){
      if (this.username != "root"){
        this.api.deleteUserAccount().subscribe((res) => {
          alert("User " + this.username + " has been deleted! Logging Out...")
          this.authService.userLogOut()
          this.router.navigate(['login']);
        })
      }else{
        alert("Cannot Delete root user!")
      }
    }
  }
}