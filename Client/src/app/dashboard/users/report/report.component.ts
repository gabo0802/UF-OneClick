import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { MatAccordion } from '@angular/material/expansion';
import { forkJoin } from 'rxjs';
import { Subscription } from 'src/app/subscription.model';
import { MatTableDataSource } from '@angular/material/table';


@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})

export class ReportComponent implements OnInit{

  @Input() username: string = '';
  @Input() subscriptionList: Subscription[] = [];
  @ViewChild(MatAccordion) accordion: MatAccordion;
  
  //Input for all of the queries
  longestSub = new MatTableDataSource<Subscription>([]);
  avgPrice = new MatTableDataSource<Subscription>([]);
  subAge = new MatTableDataSource<Subscription>([]);

  table1 : boolean = false;


  constructor(private api: ApiService, private dialogs: DialogsService) {
    this.accordion = new MatAccordion()
    this.accordion.closeAll()
  }

  ngOnInit(): void {
    
    // forkJoin({
    //   NameHere: apicallhere,
    //   AnotherNameHere: apicallhere,
    // }).subscribe({

    //   next: (res) => {
    //     //process results here, ex: variable = res.NameHere,
    //   },
    //   error: (error: HttpErrorResponse) => {
    //     //process errors here
    //   }
    // });

    this.api.subQueries(0).subscribe({

      next: (res: Subscription[]) => {

        this.subscriptionList = res;
      },
      error: (error: HttpErrorResponse) => {
        
        this.dialogs.errorDialog("ERR", "Failed to fetch query");
      }
    });
    

  }

  // Placeholder functions for the queries
  calculateLongestSub(): void {

  }

  calculateAvgPrice(): void {

  }

  calculateAvgSubAge(): void {

  }
}
