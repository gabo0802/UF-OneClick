import { HttpErrorResponse } from '@angular/common/http';
import { Component, ViewChild } from '@angular/core';
import { ValidationErrors } from '@angular/forms';
import { AbstractControl, FormControl, FormGroup, ValidatorFn, Validators } from '@angular/forms';
import { ChartConfiguration } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';

interface Month {
  name: string;
  value: number;
}

@Component({
  selector: 'app-graph',
  templateUrl: './graph.component.html',
  styleUrls: ['./graph.component.css']
})
export class GraphComponent {
  
  constructor(private api: ApiService, private dialogs: DialogsService) {}

  @ViewChild(BaseChartDirective) chart?: BaseChartDirective;

  datesForm = new FormGroup({
    'startMonth': new FormControl<number| null>(null, [Validators.required, this.datesValidator()]),
    'endMonth': new FormControl<number| null>(null, [Validators.required, this.datesValidator()]),
    'startYear': new FormControl<number| null>(null, [Validators.required, Validators.pattern('^\\d{4}$'), this.datesValidator()]),
    'endYear': new FormControl<number| null>(null, [Validators.required, Validators.pattern('^\\d{4}$'), this.datesValidator()])
  });
  
  months: Month[] = [
    { name: "January", value: 1},
    { name: "February", value: 2},
    { name: "March", value: 3},
    { name: "April", value: 4},
    { name: "May", value: 5},
    { name: "June", value: 6},
    { name: "July", value: 7},
    { name: "August", value: 8},
    { name: "September", value: 9},
    { name: "October", value: 10},
    { name: "November", value: 11},
    { name: "December", value: 12}
  ];
  
  //no ranges entered then displays nothing
  initialState: boolean = false;
  
  //ChartJs data
  barChartLegend = false;
  barChartPlugins = [];

  barChartData: ChartConfiguration<'line'>['data'] = {
    labels: [ '' ],
    datasets: [ { data: [ 0 ] } ]
  };

  barChartOptions: ChartConfiguration<'line'>['options'] = {
    responsive: true,
  };
  
  getData(): void {

    const startMonth = this.datesForm.get('startMonth')?.getRawValue();
    const startYear = parseInt(this.datesForm.get('startYear')?.getRawValue());

    const endMonth = this.datesForm.get('endMonth')?.getRawValue();
    const endYear = parseInt(this.datesForm.get('endYear')?.getRawValue());
    
    this.api.graphPrices(startMonth, startYear, endMonth, endYear).subscribe({

      next: res => {
        
        let dataStr = JSON.stringify(res);
        let data = JSON.parse(dataStr);
        let processed_data = this.processData(data);
        
        this.barChartData.datasets = [{ data: processed_data.price_data}];
        
        this.barChartData.labels = processed_data.date_labels;
        this.chart?.update();
        this.initialState = true;
      },
      error: (error: HttpErrorResponse) => {
        this.dialogs.errorDialog("Error Getting Graph!", "An error occured. Please try again later.");
      }
    })
  }

  private processData(data: {month: number, year: number, cost: number}[]): {date_labels: string[], price_data: number[]} {

    const months: string[] = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
    let date_labels: string[] = [];
    let price_data: number[] = []

    data.forEach( element => {     
      date_labels.push(months[element.month-1] + " " + element.year);
      price_data.push(element.cost);
    })
    
    return {date_labels, price_data};
  }

  //custom validator for checking date ranges
  //still has some errors, but works on a basic level
  private datesValidator(): ValidatorFn {

    return (group:AbstractControl) : ValidationErrors| null => {
  
      const startYear = group.parent?.get('startYear')?.value;
      const endYear = group.parent?.get('endYear')?.value;
      
      const startMonth = group.parent?.get('startMonth')?.value;
      const endMonth = group.parent?.get('endMonth')?.value;
      
      let status: boolean = true;
      
      if(startYear !== null && endYear !== null && startYear == endYear && startMonth > endMonth){
        status = false;        
      }
  
      if(startYear !== null && endYear !== null  && startYear > endYear){
        status = false;        
      }      
  
      return status ? null : {range: true};
    }
  }
}