import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-welcome-header',
  templateUrl: './welcome-header.component.html',
  styleUrls: ['./welcome-header.component.css']
})
export class WelcomeHeaderComponent {

  @Input() username = '';

  currentDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

}