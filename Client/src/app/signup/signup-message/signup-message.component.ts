import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';

@Component({
  selector: 'app-signup-message',
  templateUrl: './signup-message.component.html',
  styleUrls: ['./signup-message.component.css']
})
export class SignupMessageComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: {dialogTitle: string, dialogMessage: string}) { }
}
