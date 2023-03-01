import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { PasswordResetComponent } from './password-reset/password-reset.component'
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
  usernameEdit: boolean = false;
  emailEdit: boolean = false;  

  username: string = 'Loading...';
  password: string = '************';
  email: string = 'Loading...';
  timezone: string = 'Loading...';

  usernameForm = new FormGroup({
    'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
  });

  emailForm = new FormGroup({
    'email': new FormControl(this.email, [Validators.email]),
  });

  timezoneForm = new FormGroup({
    'timezonedifference': new FormControl(this.timezone, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')]),
  });

  passwordForm = new FormGroup({
    'password': new FormControl(this.password, [Validators.minLength(this.passwordCharacterLength)]),
  });

  ngOnInit(){
      this.api.getEmailandUsername().subscribe((resultMessage: string[]) => {
        this.username = resultMessage[0], 
        this.email = resultMessage[1];

        this.usernameForm = new FormGroup({
          'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
        });
      
        this.emailForm = new FormGroup({
          'email': new FormControl(this.email, [Validators.email]),
        });
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

  editUsername(): void {
    this.usernameEdit = !this.usernameEdit;
    
    if (this.usernameEdit == false){
      var oldUsername: string = this.username
      this.username = this.usernameForm.controls['username'].value || this.username

      if (this.username != oldUsername){
        if (confirm("Are you sure you want to change your username from " + oldUsername + " to " + this.username + "?")){
            if(oldUsername != "root"){
              this.api.updateUsername(this.username).subscribe((res) => {
                alert("Username Updated! Logging Out...")
                this.authService.userLogOut();
                this.router.navigate(['/login']);
              })
            }else{
              alert("Cannot Change Username of root user!")
              this.username = oldUsername;
              
              this.usernameForm = new FormGroup({
                'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
              });
            }
        }else{
          this.username = oldUsername;

          this.usernameForm = new FormGroup({
            'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
          });
        }  
      }
    }else{
      this.usernameForm = new FormGroup({
        'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
      });
    }

  }

  editEmail(): void {
    this.emailEdit = !this.emailEdit;

    if (this.emailEdit == false){
      var oldEmail: string = this.email
      this.email = this.emailForm.controls['email'].value || this.email

      if (this.email != oldEmail){
        if (confirm("Are you sure you want to change your email address from " + oldEmail + " to " + this.email + "?")){
          if(this.username != "root"){
            this.api.updateUserEmail(this.email).subscribe((res) => {
              alert("Email Updated! Please Verify Email Before Logging Back In. Logging Out...")
              this.authService.userLogOut();
              this.router.navigate(['/login']);
            })
          }else{
            alert("Cannot Change Email Address of root user!")

            this.email = oldEmail;

            this.emailForm = new FormGroup({
              'email': new FormControl(this.email, [Validators.email]),
            });
          }
        }else{
          this.email = oldEmail;

          this.emailForm = new FormGroup({
            'email': new FormControl(this.email, [Validators.email]),
          });
        }
      } 
    }else{
      this.emailForm = new FormGroup({
        'email': new FormControl(this.email, [Validators.email]),
      });
    }
  }

  editPassword(): void {
    this.callDialog();   
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

  callDialog(): void {
    this.dialog.open( PasswordResetComponent, { 
      height: '335px',
      width: '500px',
    });    
  }
}