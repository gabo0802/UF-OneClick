import { Component } from '@angular/core';

@Component({
  selector: 'app-entry',
  templateUrl: './entry.component.html',
  styleUrls: ['./entry.component.css']
})
export class EntryComponent {

  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());
}
