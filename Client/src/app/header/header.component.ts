import { Component, Input, SimpleChanges} from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})



export class HeaderComponent {  

  @Input() myHeaderState = 1;
  myNavBar: boolean= false;
  mainButtons: boolean= false;

  public HeaderComponent() {

  }

  // Changes the header state based on whether it has been changed
  ngOnChanges(changes: SimpleChanges) {
    console.log(this.myHeaderState);  
    
    switch(this.myHeaderState) {
      case 1:
        this.myNavBar = false;
        this.mainButtons = false;
        break;
      case 2:
        this.myNavBar = true;
        this.mainButtons = false;
        break;
      case 3:
        this.myNavBar = true;
        this.mainButtons = true;
        break;
      default:
        this.myNavBar = true;
        this.mainButtons = true;
    }
  }

}


