import { Component } from '@angular/core';
import { ApiService } from 'src/app/api.service';

@Component({
  selector: 'app-reset',
  templateUrl: './reset.component.html',
  styleUrls: ['./reset.component.css']
})
export class ResetComponent {
  constructor(private api: ApiService) { }
  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

  adminReset(): void{
    if (confirm("All Data On UF-OneClick Will Be Reset!")){
      this.api.resetWebsite().subscribe((res) => {
          alert("All Data On UF-OneClick Has Been Reset!");
      });
    }
  }
}
