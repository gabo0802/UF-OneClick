import { HttpErrorResponse } from '@angular/common/http';
import { Component, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { MatAccordion } from '@angular/material/expansion';


@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})

export class ReportComponent {
  @ViewChild(MatAccordion) accordion: MatAccordion;

  constructor(private api: ApiService, private dialogs: DialogsService) {
    this.accordion = new MatAccordion()
    this.accordion.openAll()
  }

}
