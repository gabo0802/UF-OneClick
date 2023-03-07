import { Component, OnInit} from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api.service';
import { AuthService } from '../auth.service';
import { DialogsService } from '../dialogs.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit{

  constructor(private router: Router, private api: ApiService, private authService: AuthService, private dialogs: DialogsService) {}

  hide: boolean = true;  

  username: string = '';
  email: string = '';
  timezone: string = '';  

  ngOnInit(){  

    this.api.getEmailandUsername().subscribe((res: Object) => {

      let data = JSON.stringify(res);
      let userInfo = JSON.parse(data);
      this.username = userInfo.username;          
    });

    this.api.getTimezone().subscribe((resultMessage: string[]) => {
      this.timezone = resultMessage[1] + "UTC";      
      
    });

  }
 
  back(): void {
    this.router.navigate(['users']);
  }

  delete(): void {

    if(this.username !== 'root'){
      this.dialogs.deleteAccount();
    }        
  }
}