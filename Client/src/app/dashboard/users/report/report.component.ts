import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { MatAccordion } from '@angular/material/expansion';
import { forkJoin } from 'rxjs';


@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})

export class ReportComponent implements OnInit{

  @Input() username: string = '';

  @ViewChild(MatAccordion) accordion: MatAccordion;

  constructor(private api: ApiService, private dialogs: DialogsService) {
    this.accordion = new MatAccordion()
    this.accordion.openAll()
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
  }

}
