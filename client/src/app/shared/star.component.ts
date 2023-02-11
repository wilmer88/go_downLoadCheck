import {Component,Input, OnChanges} from "@angular/core";


@Component({
    selector: "happy-stars",
    templateUrl: "./star.component.html",
    styleUrls:["./star.component.css"]
})


export class StarComponent{
   @Input() rating: number = 0;
    cropWidth: number = 75;
    
    ngOnChanges(): void {
        this.cropWidth= this.rating * 75/5;
    }

    onClick(): void {
        console.log(`the ${this.rating} star rating was clicked`)

    }

};