import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';

@Component({
  selector: 'app-login-message',
  templateUrl: './login-message.component.html',
  styleUrls: ['./login-message.component.css']
})
export class LoginMessageComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: {dialogTitle: string, dialogMessage: string}) { }

}
