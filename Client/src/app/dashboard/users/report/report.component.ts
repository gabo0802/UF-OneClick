import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { MatAccordion } from '@angular/material/expansion';
import { forkJoin } from 'rxjs';
import { Subscription } from 'src/app/subscription.model';
import { MatTableDataSource } from '@angular/material/table';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})

export class ReportComponent implements OnInit{

  @Input() username: String = '';
  cost: String = '$0.00';
  @Input() subscriptionList: Subscription[] = [];
  @ViewChild(MatAccordion) accordion: MatAccordion;
  
  //Input for all of the queries
  myQueries : String[] = [];
  panelOpenState : boolean = true;
  comparePriceForm: FormGroup = {} as FormGroup;

  constructor(private api: ApiService, private dialogs: DialogsService) {
    this.accordion = new MatAccordion()
    this.accordion.openAll()
  }

  ngOnInit(): void {    

    // This is randomly putting the queries into my array, not sure why
    for(let i = 0; i < 8; i++) {
      this.api.subQueries(i).subscribe({

        next: (res: String[]) => {
  
          this.myQueries = this.myQueries.concat(res)
          console.log(this.myQueries)
        },
        error: (error: HttpErrorResponse) => {
          
          this.dialogs.errorDialog("ERR", "Failed to fetch query");
        }
      });
    }

    this.comparePriceForm = new FormGroup({
      'month': new FormControl(null, [Validators.required, Validators.pattern('^[0-9]{2}$')]),
      'year': new FormControl(null, [Validators.required, Validators.pattern('^[0-9]{4}$')]),
    });
  }

  onSubmit(){
    var monthString:string = this.comparePriceForm.get('month')?.value;
    var yearString:string = this.comparePriceForm.get('year')?.value;

    console.log(monthString)
    console.log(yearString)

    var monthNumber: number = +monthString
    var yearNumber: number = +yearString

    console.log(yearNumber)
    console.log(monthNumber)

    this.api.comparePrice(monthNumber, yearNumber).subscribe({  
      next: (res) => {
        this.cost = JSON.stringify(res);
      }
    })
  }

  // Order of queries (for reference purposes), each query outputs a string array of length 2, the query title and output /  output name and description.
  /*
      case 0:
        URL = "/api/longestsub";
        break;
      case 1:
        URL = "/api/longestcontinoussub";
        break;
      case 2:
        URL = "/api/longestactivesub";
        break;
      case 3:
        URL = "/api/avgpriceactivesub";
        break;
      case 4:
        URL = "/api/avgpriceallsubs";
        break;
      case 5:
        URL = "/api/avgageallsubs";
        break;
      case 6:
        URL = "/api/avgageactivesubs";
        break;
      case 7:
        URL = "/api/avgagecontinuoussubs";
        break;
  */
}
