import { Component } from '@angular/core';
import { ApiService } from 'src/app/api.service';

@Component({
  selector: 'app-newsletter',
  templateUrl: './newsletter.component.html',
  styleUrls: ['./newsletter.component.css']
})
export class NewsletterComponent {
  
  constructor(private api: ApiService) { }
  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

  newsLetter(): void{
    var message:string = (document.getElementById('newslettermessage') as HTMLInputElement).value;

    if(message != "Enter Message For Newsletter!"){
      if(confirm("Message Will Be Sent To All Users!")){
        this.api.sendNews(message).subscribe( (res) => {
            alert("Message Sent");
            (document.getElementById('newslettermessage') as HTMLInputElement).value = "Enter Message For Newsletter!";
        })
      }
    }
  }

  resetMessage(): void{
    (document.getElementById('newslettermessage') as HTMLInputElement).value = "Enter Message For Newsletter!";
  }

  remindTime(): void{
    this.api.sendDailyRemind().subscribe((res) => {
      alert(JSON.stringify(res).replaceAll("\"","").replace("{","").replace("}","").replace(":",": "));
    });
  }
}
