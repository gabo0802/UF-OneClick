import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-newsletter',
  templateUrl: './newsletter.component.html',
  styleUrls: ['./newsletter.component.css']
})
export class NewsletterComponent {
  
  constructor(private api: ApiService, private dialogs: DialogsService) { }

  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

  newsletter: FormControl = new FormControl<string>('', [Validators.required, Validators.pattern('^.*\\S.*[a-zA-z0-9 ]*$')]);

  sendNewsletter(): void{
    let message: string = this.newsletter.getRawValue();
    
    if(message !== ' '){

      if(confirm("Message will be sent to all users!")){

        this.api.sendNews(message).subscribe({

          next: (res) => {

            this.dialogs.successDialog("Message sent to all users.");
            this.newsletter.reset();
          },
          error: (error: HttpErrorResponse) => {
            this.dialogs.errorDialog("Error sending Message", "An error occured please try again later.");
          }
        })
      }
    }
  }

  resetMessage(): void{
    this.newsletter.reset();
  }

  remindTime(): void{
    this.api.sendDailyRemind().subscribe((res) => {
      alert(JSON.stringify(res).replaceAll("\"","").replace("{","").replace("}","").replace(":",": "));
    });
  }
}
