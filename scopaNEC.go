package main//esse é idêntico ao scopatroca.go, mas não se escolhe quem começa e tem um robô nível 5

import (
	"fmt"
	"math"
	"math/rand"
	"time"//a fim de que o jogador possa ler o resultado
)

type carta struct{
    naipe string;
    valor int;//1, 2, 3, 4, 5, 6, 7, 8(fante), 9(cavalo) ,10(rei)
}

func novacarta(naipe string, valor int) carta{
	c:=carta{naipe:naipe, valor:valor};
	return c;
}

func roboescolha(mesa []carta, mao []carta, nivel int)int {
	var pontcarta []int;
	for i:=0; i<len(mao); i++{
		pontcarta=append(pontcarta, 0);//tentativa de fazer um robô melhor do que o 2(seria trocar esse 0 por -10)
	}
	var n int=0;//qual carta
	for i:=0; i<len(pontcarta); i++{
		pontcarta[i]=0;
	}
	if nivel==0{//tenta sempre deitar
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]-=1;
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]-=2;
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]-=3;
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if pontcarta[i]>pontcarta[n]{
				n=i;
			}
		}
		return n;
	}else if nivel==1{//aleatório
		return rand.Intn(len(mao));
	}else if nivel==2{//pega a primeira
		return 0;
	}else if nivel==3{//tenta sempre capturar
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]+=1;
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]+=2;
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]+=3;
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if pontcarta[i]>pontcarta[n]{
				n=i;
			}
		}
		return n;
	}else if nivel==4{//tenta sempre capturar (com uma reviravolta?)
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]-=1;
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]+=2;
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]+=3;
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if math.Abs(float64(pontcarta[i]))>math.Abs(float64(pontcarta[n])){
				n=i;
			}
		}
		return n;
	}else{//tenta sempre capturar, mas é ganancioso
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]+=1;
					if mesa[j].naipe=="denara"||mesa[j].valor==7{
						if len(mesa)-1==0{//se for escopa
							return i;
						}
						pontcarta[i]+=1;
					}
					if len(mesa)-1==0{
						pontcarta[i]+=1;
					}
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]+=2;
						if (mesa[j].naipe=="denara"||mesa[j].valor==7)||(mesa[k].naipe=="denara"||mesa[k].valor==7){
							if len(mesa)-2==0{//se for escopa
								return i;
							}
							pontcarta[i]+=2;
						}
						if len(mesa)-2==0{
							pontcarta[i]+=2;
						}
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]+=3;
							if (mesa[j].naipe=="denara"||mesa[j].valor==7)||(mesa[k].naipe=="denara"||mesa[k].valor==7)||(mesa[h].naipe=="denara"||mesa[h].valor==7){
								if len(mesa)-3==0{//se for escopa
									return i;
								}
								pontcarta[i]+=3;
							}
							if len(mesa)-3==0{
								pontcarta[i]+=3;
							}
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if math.Abs(float64(pontcarta[i]))>math.Abs(float64(pontcarta[n])){
				n=i;
			}
		}
		return n;
	}
}

func pontosmn(a int, b int) int{
	if b==7{
		return b;
	}
	//a partir daqui b!=7
	if a==7{
		return a;
	}else if a==6{
		return a;
	}else if a==1&&(((b<6)&&(b>1))||b>7){
		return a;
	}else if a==5&&(((b<6)&&(b>1))||b>7){
		return a;
	}else if a==4&&(((b<5)&&(b>1))||b>7){
		return a;
	}else if a==3&&(((b<4)&&(b>1))||b>7){
		return a;
	}else if a==2&&(b==2||b>7){
		return a;
	}else{
		return b;
	}
}
func main(){
	//deleta:="\033[2J\033[H";// esse deleta funciona no terminal do VS code, mas não funciona quando compilo
	var vez int=0;//se par, vez do jogador; se impar, vez do computador 
	var baralho [] carta;//vai ter 40
	var naipes = [4] string{"denara", "paus", "espadas", "copas"};
	var a int=-1;
	//cria o baralho
	for i:=0; i<4; i++{
		for j:=1; j<11; j++{
			a++;
			baralho=append(baralho, novacarta(naipes[i],j));
		}
	}
	/*
	for i:=0; i<len(baralho); i++{
		fmt.Println(baralho[i]);
	}
	*/

	//embaralha o baralho
	genesis:
	var troca carta;
	var b int;
	var c int;
	for i:=0; i<80; i++{
		b=rand.Intn(40);
		c=rand.Intn(40);
		troca.valor=baralho[b].valor;
		troca.naipe=baralho[b].naipe;
		baralho[b].valor=baralho[c].valor;
		baralho[b].naipe=baralho[c].naipe;
		baralho[c].valor=troca.valor;
		baralho[c].naipe=troca.naipe;
	}
	//cria as mãos e mesa e as preenche
	var mesa[]carta;
	var maojogador[]carta;//cada mão tem no máximo, três cartas
	var maocomputador[]carta;
	var escopasjogador int;
	var escopascomputador int;
	var montejogador[]carta;//conte o monte
	var montecomputador[]carta;
	var escolhajogador int;
	var escolhacomputador int;

	mesa=baralho[:4];
	baralho=baralho[4:];
	qtdreis:=0;
	for i:=0; i<len(mesa); i++{
		if mesa[i].valor==10{
			qtdreis++;
		}
	}
	if qtdreis>=3{
		goto genesis;
	}
	nivelcomputador:=0;

	vez=rand.Intn(2);
	fmt.Println("Contra que robô você quer jogar?[0 a 2 - fácil, 3 - médio, 4 e 5-difícil?]");
	fmt.Scanln(&nivelcomputador);
	for len(baralho)>=0{//enquanto houver cartas distribuíveis(não nulas)
		distribuicartas:
			if len(baralho)==0{
				if (vez-1)%2==0{
					for i:=0; i<len(mesa); i++{
						montejogador=append(montejogador, mesa[i]);
					}
					goto contagem;
				}else{
					for i:=0; i<len(mesa); i++{
						montecomputador=append(montecomputador, mesa[i]);
					}
					goto contagem;
				}
			}
			//distribui novo com for
			for i:=0; i<3; i++{
				maojogador=append(maojogador, baralho[i]);
			}
			//maojogador=baralho[:3];//pensei em usar um append, mas nâo é necessário; se for, fazer devidas alterações
			baralho=baralho[3:];
			for i:=0; i<3; i++{
				maocomputador=append(maocomputador, baralho[i]);
			}
			//maocomputador=baralho[:3];
			baralho=baralho[3:];
		for (len(maojogador)>0&&len(maocomputador)>=0)||(len(maojogador)>=0&&len(maocomputador)>0){
			l:=0;
			for i:=0; i<len(mesa); i++{
				if mesa[i].naipe=="denara"{
					troca.valor=mesa[i].valor;
					troca.naipe=mesa[i].naipe;
					mesa[i].valor=mesa[l].valor;
					mesa[i].naipe=mesa[l].naipe;
					mesa[l].valor=troca.valor;
					mesa[l].naipe=troca.naipe;
					l++;
				}
			}
			//jogo
			if vez%2==0{
			//imprime as cartas
			//fmt.Print(string(deleta));//ver se funciona
			fmt.Println("\n");
			fmt.Println("\n");
			fmt.Print("		");
			for i:=0; i<len(maocomputador); i++{
				fmt.Print("[] ");
			}
			fmt.Println("\n");
			for i:=0; i<len(mesa); i++{
				fmt.Print("[");
				fmt.Print(mesa[i].valor);
				fmt.Print(" de ");
				fmt.Print(mesa[i].naipe);
				fmt.Print("]   ");
			}
			//fmt.Println(mesa);
			fmt.Println("");
			fmt.Println("");
			//fmt.Println(maojogador);
			fmt.Print("   ");
			for i:=0; i<len(maojogador); i++{
				fmt.Print(i);
				fmt.Print(":");
				fmt.Print("[");
				fmt.Print(maojogador[i].valor);
				fmt.Print(" de ");
				fmt.Print(maojogador[i].naipe);
				fmt.Print("]");
				fmt.Print("  ");
			}
			fmt.Println("\n");
				coletadedados:
				fmt.Scanln(&escolhajogador);
				if escolhajogador>=len(maojogador)||escolhajogador<0{
					goto coletadedados;
				}
				//fmt.Print("tu:");
				//fmt.Println(escolhajogador);
				for i:=0; i<len(mesa); i++{//se houver uma carta de valor identico na mesa
					if mesa[i].valor==maojogador[escolhajogador].valor{
						if len(mesa)-1==0{
							escopasjogador++;
							fmt.Println("Escopa do jogador!");
						}
						troca.valor=maojogador[escolhajogador].valor;
						troca.naipe=maojogador[escolhajogador].naipe;
						maojogador[escolhajogador].valor=maojogador[0].valor;
						maojogador[escolhajogador].naipe=maojogador[0].naipe;
						maojogador[0].valor=troca.valor;//coloca carta no começo pela praticidade
						maojogador[0].naipe=troca.naipe;
						//coloca cartas nos devidos lugares
						montejogador=append(montejogador, maojogador[0]);
						montejogador=append(montejogador, mesa[i]);
						maojogador=maojogador[1:];
						troca.valor=mesa[i].valor;
						troca.naipe=mesa[i].naipe;
						mesa[i].valor=mesa[0].valor;
						mesa[i].naipe=mesa[0].naipe;
						mesa[0].valor=troca.valor;
						mesa[0].naipe=troca.naipe;
						mesa=mesa[1:];
						
						goto vez;
					}
				}
				if len(mesa)>=4{//só coloquei, porque joguei e o computador não contou minha escopa de quatro
					for i:=0; i<len(mesa)-3; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-2; j++{
							for k:=j+1; k<len(mesa)-1; k++{
								for g:=k+1; g<len(mesa); g++{
									if mesa[i].valor+mesa[j].valor+mesa[k].valor+mesa[g].valor==maojogador[escolhajogador].valor{
										if len(mesa)-4==0{
											escopasjogador++;
											fmt.Println("Escopa do jogador!");
										}
										//fmt.Println(i);fmt.Println(j);
										troca.valor=mesa[i].valor;
										troca.naipe=mesa[i].naipe;
										mesa[i].valor=mesa[0].valor;
										mesa[i].naipe=mesa[0].naipe;
										mesa[0].valor=troca.valor;
										mesa[0].naipe=troca.naipe;
										montejogador=append(montejogador, mesa[0]);
										//mesa=mesa[1:];//carta 1 da mesa
										troca.valor=mesa[j].valor;
										troca.naipe=mesa[j].naipe;
										mesa[j].valor=mesa[1].valor;
										mesa[j].naipe=mesa[1].naipe;
										mesa[1].valor=troca.valor;
										mesa[1].naipe=troca.naipe;
										montejogador=append(montejogador, mesa[1]);
										//
										troca.valor=mesa[k].valor;
										troca.naipe=mesa[k].naipe;
										mesa[k].valor=mesa[2].valor;
										mesa[k].naipe=mesa[2].naipe;
										mesa[2].valor=troca.valor;
										mesa[2].naipe=troca.naipe;
										montejogador=append(montejogador, mesa[2]);
										//
										troca.valor=mesa[g].valor;
										troca.naipe=mesa[g].naipe;
										mesa[g].valor=mesa[3].valor;
										mesa[g].naipe=mesa[3].naipe;
										mesa[3].valor=troca.valor;
										mesa[3].naipe=troca.naipe;
										montejogador=append(montejogador, mesa[3]);
				
										mesa=mesa[4:];//carta 4 da mesa
										//fmt.Println("!!!!!!!!!!!!!!!!");
										troca.valor=maojogador[escolhajogador].valor;
										troca.naipe=maojogador[escolhajogador].naipe;
										maojogador[escolhajogador].valor=maojogador[0].valor;
										maojogador[escolhajogador].naipe=maojogador[0].naipe;
										maojogador[0].valor=troca.valor;
										maojogador[0].naipe=troca.naipe;
										montejogador=append(montejogador, maojogador[0]);
										maojogador=maojogador[1:];
										goto vez;
									}
								}
							}
						}
					}
				}
				if len(mesa)>=3{
					for i:=0; i<len(mesa)-2; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-1; j++{
							for k:=j+1; k<len(mesa); k++{
								if mesa[i].valor+mesa[j].valor+mesa[k].valor==maojogador[escolhajogador].valor{
									if len(mesa)-3==0{
										escopasjogador++;
										fmt.Println("Escopa do jogador!");
									}
									troca.valor=mesa[i].valor;
									troca.naipe=mesa[i].naipe;
									mesa[i].valor=mesa[0].valor;
									mesa[i].naipe=mesa[0].naipe;
									mesa[0].valor=troca.valor;
									mesa[0].naipe=troca.naipe;
									montejogador=append(montejogador, mesa[0]);
									//mesa=mesa[1:];//carta 1 da mesa
									troca.valor=mesa[j].valor;
									troca.naipe=mesa[j].naipe;
									mesa[j].valor=mesa[1].valor;
									mesa[j].naipe=mesa[1].naipe;
									mesa[1].valor=troca.valor;
									mesa[1].naipe=troca.naipe;
									montejogador=append(montejogador, mesa[1]);
									//
									troca.valor=mesa[k].valor;
									troca.naipe=mesa[k].naipe;
									mesa[k].valor=mesa[2].valor;
									mesa[k].naipe=mesa[2].naipe;
									mesa[2].valor=troca.valor;
									mesa[2].naipe=troca.naipe;
									montejogador=append(montejogador, mesa[2]);
			
									mesa=mesa[3:];//carta 4 da mesa
									//fmt.Println("!!!!!!!!!!!!!!!!");
									troca.valor=maojogador[escolhajogador].valor;
									troca.naipe=maojogador[escolhajogador].naipe;
									maojogador[escolhajogador].valor=maojogador[0].valor;
									maojogador[escolhajogador].naipe=maojogador[0].naipe;
									maojogador[0].valor=troca.valor;
									maojogador[0].naipe=troca.naipe;

									montejogador=append(montejogador, maojogador[0]);
									maojogador=maojogador[1:];
									goto vez;
								}
							}
						}
					}
				}
				for i:=0; i<len(mesa)-1; i++{//se houver mais de uma carta que somada dá a escolha
					for j:=i+1; j<len(mesa); j++{
						if mesa[i].valor+mesa[j].valor==maojogador[escolhajogador].valor{
							if len(mesa)-2==0{
								escopasjogador++;
								fmt.Println("Escopa do jogador!");
							}
							//fmt.Println(i);fmt.Println(j);
							troca.valor=mesa[i].valor;
							troca.naipe=mesa[i].naipe;
							mesa[i].valor=mesa[0].valor;
							mesa[i].naipe=mesa[0].naipe;
							mesa[0].valor=troca.valor;
							mesa[0].naipe=troca.naipe;
							montejogador=append(montejogador, mesa[0]);
							//mesa=mesa[1:];//carta 1 da mesa
							troca.valor=mesa[j].valor;
							troca.naipe=mesa[j].naipe;
							mesa[j].valor=mesa[1].valor;
							mesa[j].naipe=mesa[1].naipe;
							mesa[1].valor=troca.valor;
							mesa[1].naipe=troca.naipe;
							montejogador=append(montejogador, mesa[1]);
							mesa=mesa[2:];//carta 2 da mesa
							//fmt.Println("!!!!!!!!!!!!!!!!");
							troca.valor=maojogador[escolhajogador].valor;
							troca.naipe=maojogador[escolhajogador].naipe;
							maojogador[escolhajogador].valor=maojogador[0].valor;
							maojogador[escolhajogador].naipe=maojogador[0].naipe;
							maojogador[0].valor=troca.valor;
							maojogador[0].naipe=troca.naipe;
							montejogador=append(montejogador, maojogador[0]);
							maojogador=maojogador[1:];
							goto vez;
						}
					}
				}
				
				//se não houver somas de cartas na mesa que se igualem à carta, a carta vai à mesa
				troca.valor=maojogador[escolhajogador].valor;
				troca.naipe=maojogador[escolhajogador].naipe;
				maojogador[escolhajogador].valor=maojogador[0].valor;
				maojogador[escolhajogador].naipe=maojogador[0].naipe;
				maojogador[0].valor=troca.valor;
				maojogador[0].naipe=troca.naipe;//coloca carta no começo pela praticidade
				mesa=append(mesa, maojogador[0]);
				maojogador=maojogador[1:];
				
				vez:
				vez++;
				//jogador
			}else{//fazer algoritmo inteligente
				//algoritmo computador burro
				//escolhacomputador=rand.Intn(len(maocomputador));
				escolhacomputador=roboescolha(mesa, maocomputador, nivelcomputador);
				for i:=0; i<len(mesa); i++{//se houver uma carta de valor identico na mesa
					if mesa[i].valor==maocomputador[escolhacomputador].valor{
						if len(mesa)-1==0{
							escopascomputador++;
							fmt.Println("Escopa do computador!");
						}
						troca.valor=maocomputador[escolhacomputador].valor;
						troca.naipe=maocomputador[escolhacomputador].naipe;
						maocomputador[escolhacomputador].valor=maocomputador[0].valor;
						maocomputador[escolhacomputador].naipe=maocomputador[0].naipe;
						maocomputador[0].valor=troca.valor;//coloca carta no começo pela praticidade
						maocomputador[0].naipe=troca.naipe;
						//coloca cartas nos devidos lugares
						montecomputador=append(montecomputador, maocomputador[0]);
						montecomputador=append(montecomputador, mesa[i]);
						maocomputador=maocomputador[1:];
						troca.valor=mesa[i].valor;
						troca.naipe=mesa[i].naipe;
						mesa[i].valor=mesa[0].valor;
						mesa[i].naipe=mesa[0].naipe;
						mesa[0].valor=troca.valor;
						mesa[0].naipe=troca.naipe;
						mesa=mesa[1:];
						
						goto vezc;
					}
				}
				if len(mesa)>=4{//só coloquei, porque joguei e o computador não contou minha escopa de quatro
					for i:=0; i<len(mesa)-3; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-2; j++{
							for k:=j+1; k<len(mesa)-1; k++{
								for g:=k+1; g<len(mesa); g++{
									if mesa[i].valor+mesa[j].valor+mesa[k].valor+mesa[g].valor==maocomputador[escolhacomputador].valor{
										if len(mesa)-4==0{
											escopascomputador++;
											fmt.Println("Escopa do computador!");
										}
										//fmt.Println(i);fmt.Println(j);
										troca.valor=mesa[i].valor;
										troca.naipe=mesa[i].naipe;
										mesa[i].valor=mesa[0].valor;
										mesa[i].naipe=mesa[0].naipe;
										mesa[0].valor=troca.valor;
										mesa[0].naipe=troca.naipe;
										montecomputador=append(montecomputador, mesa[0]);
										//mesa=mesa[1:];//carta 1 da mesa
										troca.valor=mesa[j].valor;
										troca.naipe=mesa[j].naipe;
										mesa[j].valor=mesa[1].valor;
										mesa[j].naipe=mesa[1].naipe;
										mesa[1].valor=troca.valor;
										mesa[1].naipe=troca.naipe;
										montecomputador=append(montecomputador, mesa[1]);
										//
										troca.valor=mesa[k].valor;
										troca.naipe=mesa[k].naipe;
										mesa[k].valor=mesa[2].valor;
										mesa[k].naipe=mesa[2].naipe;
										mesa[2].valor=troca.valor;
										mesa[2].naipe=troca.naipe;
										montecomputador=append(montecomputador, mesa[2]);
										//
										troca.valor=mesa[g].valor;
										troca.naipe=mesa[g].naipe;
										mesa[g].valor=mesa[3].valor;
										mesa[g].naipe=mesa[3].naipe;
										mesa[3].valor=troca.valor;
										mesa[3].naipe=troca.naipe;
										montecomputador=append(montecomputador, mesa[3]);

										mesa=mesa[4:];//carta 4 da mesa
										//fmt.Println("!!!!!!!!!!!!!!!!");
										troca.valor=maocomputador[escolhacomputador].valor;
										troca.naipe=maocomputador[escolhacomputador].naipe;
										maocomputador[escolhacomputador].valor=maocomputador[0].valor;
										maocomputador[escolhacomputador].naipe=maocomputador[0].naipe;
										maocomputador[0].valor=troca.valor;
										maocomputador[0].naipe=troca.naipe;
										montecomputador=append(montecomputador, maocomputador[0]);
										maocomputador=maocomputador[1:];
										goto vezc;
									}
								}
							}
						}
					}
				}
				if len(mesa)>=3{
					for i:=0; i<len(mesa)-2; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-1; j++{
							for k:=j+1; k<len(mesa); k++{
								if mesa[i].valor+mesa[j].valor+mesa[k].valor==maocomputador[escolhacomputador].valor{
									if len(mesa)-3==0{
										escopascomputador++;
										fmt.Println("Escopa do computador!");
									}
									troca.valor=mesa[i].valor;
									troca.naipe=mesa[i].naipe;
									mesa[i].valor=mesa[0].valor;
									mesa[i].naipe=mesa[0].naipe;
									mesa[0].valor=troca.valor;
									mesa[0].naipe=troca.naipe;
									montecomputador=append(montecomputador, mesa[0]);
									//mesa=mesa[1:];//carta 1 da mesa
									troca.valor=mesa[j].valor;
									troca.naipe=mesa[j].naipe;
									mesa[j].valor=mesa[1].valor;
									mesa[j].naipe=mesa[1].naipe;
									mesa[1].valor=troca.valor;
									mesa[1].naipe=troca.naipe;
									montecomputador=append(montecomputador, mesa[1]);
									//
									troca.valor=mesa[k].valor;
									troca.naipe=mesa[k].naipe;
									mesa[k].valor=mesa[2].valor;
									mesa[k].naipe=mesa[2].naipe;
									mesa[2].valor=troca.valor;
									mesa[2].naipe=troca.naipe;
									montecomputador=append(montecomputador, mesa[2]);

									mesa=mesa[3:];//carta 4 da mesa
									//fmt.Println("!!!!!!!!!!!!!!!!");
									troca.valor=maocomputador[escolhacomputador].valor;
									troca.naipe=maocomputador[escolhacomputador].naipe;
									maocomputador[escolhacomputador].valor=maocomputador[0].valor;
									maocomputador[escolhacomputador].naipe=maocomputador[0].naipe;
									maocomputador[0].valor=troca.valor;
									maocomputador[0].naipe=troca.naipe;

									montecomputador=append(montecomputador, maocomputador[0]);
									maocomputador=maocomputador[1:];
									goto vezc;
								}
							}
						}
					}
				}
				for i:=0; i<len(mesa)-1; i++{//se houver mais de uma carta que somada dá a escolha
					for j:=i+1; j<len(mesa); j++{
						if mesa[i].valor+mesa[j].valor==maocomputador[escolhacomputador].valor{
							if len(mesa)-2==0{
								escopascomputador++;
								fmt.Println("Escopa do computador!");
							}
							//fmt.Println(i);fmt.Println(j);
							troca.valor=mesa[i].valor;
							troca.naipe=mesa[i].naipe;
							mesa[i].valor=mesa[0].valor;
							mesa[i].naipe=mesa[0].naipe;
							mesa[0].valor=troca.valor;
							mesa[0].naipe=troca.naipe;
							montecomputador=append(montecomputador, mesa[0]);
							//mesa=mesa[1:];//carta 1 da mesa
							troca.valor=mesa[j].valor;
							troca.naipe=mesa[j].naipe;
							mesa[j].valor=mesa[1].valor;
							mesa[j].naipe=mesa[1].naipe;
							mesa[1].valor=troca.valor;
							mesa[1].naipe=troca.naipe;
							montecomputador=append(montecomputador, mesa[1]);
							mesa=mesa[2:];//carta 2 da mesa
							//fmt.Println("!!!!!!!!!!!!!!!!");
							troca.valor=maocomputador[escolhacomputador].valor;
							troca.naipe=maocomputador[escolhacomputador].naipe;
							maocomputador[escolhacomputador].valor=maocomputador[0].valor;
							maocomputador[escolhacomputador].naipe=maocomputador[0].naipe;
							maocomputador[0].valor=troca.valor;
							maocomputador[0].naipe=troca.naipe;
							montecomputador=append(montecomputador, maocomputador[0]);
							maocomputador=maocomputador[1:];
							goto vezc;
						}
					}
				}

				//se não houver somas de cartas na mesa que se igualem à carta, a carta vai à mesa
				troca.valor=maocomputador[escolhacomputador].valor;
				troca.naipe=maocomputador[escolhacomputador].naipe;
				maocomputador[escolhacomputador].valor=maocomputador[0].valor;
				maocomputador[escolhacomputador].naipe=maocomputador[0].naipe;
				maocomputador[0].valor=troca.valor;
				maocomputador[0].naipe=troca.naipe;//coloca carta no começo pela praticidade
				mesa=append(mesa, maocomputador[0]);
				maocomputador=maocomputador[1:];
				vezc:
				vez++;
			}
		}
		goto distribuicartas;//quando acabarem as cartas, distribui-las é o correto.
	}
	contagem:
	pontosj:=0;
	pontosc:=0;
	setebelo:=false;//true-jogador; false-computador
	ouroj:=0;
	ouroc:=0;
	fmt.Print("Jogador: ");
	fmt.Print(len(montejogador));
	fmt.Println(" cartas.");
	fmt.Print("Computador: ");
	fmt.Print(len(montecomputador));
	fmt.Println(" cartas.");
	//só falar escopas se  ocorreram
	if escopasjogador!=0{
		if escopasjogador>1{
			fmt.Print("Jogador: ");
			fmt.Print(escopasjogador);
			fmt.Println(" escopas.");
		}else{
			fmt.Print("Jogador: ");
			fmt.Print(escopasjogador);
			fmt.Println(" escopa.");
		}
	}
	if escopascomputador!=0{
		if escopascomputador>1{
			fmt.Print("Computador: ");
			fmt.Print(escopascomputador);
			fmt.Println(" escopas.");
		}else{
			fmt.Print("Computador: ");
			fmt.Print(escopascomputador);
			fmt.Println(" escopa.");
		}
	}
	if len(montejogador)>len(montecomputador){
		pontosj++;
	}else if len(montejogador)<len(montecomputador){
		pontosc++;
	}else{
		pontosj+=0;
	}
	//Quem capturou o sete de ouros
	for i:=0; i<len(montejogador); i++{
		if montejogador[i].naipe=="denara"&&montejogador[i].valor==7{
			fmt.Println("Jogador capturou o sete de ouros.");
			setebelo=true;
		}
	}
	if setebelo{
		pontosj++;
	}else{
		fmt.Println("Computador capturou o sete de ouros.");
		pontosc++;
	}
	//fmt.Println("Quem capturou mais cartas do naipe de ouros?");
	for i:=0; i<len(montejogador); i++{
		if montejogador[i].naipe=="denara"{
			ouroj++;
		}
	}
	for i:=0; i<len(montecomputador); i++{
		if montecomputador[i].naipe=="denara"{
			ouroc++;
		}
	}
	if ouroj>ouroc{
		pontosj++;
		fmt.Print("Jogador: ");
		fmt.Print(ouroj);
		fmt.Println(" cartas do naipe de ouros");
	}else if ouroc>ouroj{
		pontosc++;
		fmt.Print("Computador: ");
		fmt.Print(ouroc);
		fmt.Println(" cartas do naipe de ouros");
	}else{
		pontosj+=0;
	}
	//fmt.Println("Quem teve o melhor número?");//"primiera" ou algo do tipo
	mnj:=[4]int {0, 0, 0, 0};//ouro, pau, espada, copa 
	mnc:=[4]int {0, 0, 0, 0};
	for i:=0; i<len(montejogador); i++{
		switch(montejogador[i].naipe){
			case "denara":
				if pontosmn(montejogador[i].valor, mnj[0])==montejogador[i].valor{//montejogador[i].valor>mnj[0]{
					mnj[0]=montejogador[i].valor;
				}
			case "paus":
				if pontosmn(montejogador[i].valor, mnj[1])==montejogador[i].valor{//montejogador[i].valor>mnj[1]{
					mnj[1]=montejogador[i].valor;
				}
			case "espadas":
				if pontosmn(montejogador[i].valor, mnj[2])==montejogador[i].valor{//montejogador[i].valor>mnj[2]{
					mnj[2]=montejogador[i].valor;
				}
			case "copas":
				if pontosmn(montejogador[i].valor, mnj[3])==montejogador[i].valor{//montejogador[i].valor>mnj[3]{
					mnj[3]=montejogador[i].valor;
				}
		}
	}
	for i:=0; i<len(montecomputador); i++{
		switch(montecomputador[i].naipe){
			case "denara":
				if pontosmn(montecomputador[i].valor, mnj[0])==montecomputador[i].valor{//montecomputador[i].valor>mnj[0]{
					mnj[0]=montecomputador[i].valor;
				}
			case "paus":
				if pontosmn(montecomputador[i].valor, mnj[1])==montecomputador[i].valor{//montecomputador[i].valor>mnj[1]{
					mnj[1]=montecomputador[i].valor;
				}
			case "espadas":
				if pontosmn(montecomputador[i].valor, mnj[2])==montecomputador[i].valor{//montecomputador[i].valor>mnj[2]{
					mnj[2]=montecomputador[i].valor;
				}
			case "copas":
				if pontosmn(montecomputador[i].valor, mnj[3])==montecomputador[i].valor{//montecomputador[i].valor>mnj[3]{
					mnj[3]=montecomputador[i].valor;
				}
		}
	}
	for i:=0; i<len(mnj); i++{//os dois fors são obrigatórios (consulte a wikipédia)
		switch(mnj[i]){
		case 7:
			mnj[i]=21;
		case 6:
			mnj[i]=18;
		case 1:
			mnj[i]=16;
		case 5:
			mnj[i]=15;
		case 4:
			mnj[i]=14;
		case 3:
			mnj[i]=13;
		case 2:
			mnj[i]=12;
		default:
			mnj[i]=10;
		}
	}
	for i:=0; i<len(mnc); i++{
		switch(mnc[i]){
		case 7:
			mnc[i]=21;
		case 6:
			mnc[i]=18;
		case 1:
			mnc[i]=16;
		case 5:
			mnc[i]=15;
		case 4:
			mnc[i]=14;
		case 3:
			mnc[i]=13;
		case 2:
			mnc[i]=12;
		default:
			mnc[i]=10;
		}
	}
	melhumeroc:=mnc[0]+mnc[1]+mnc[2]+mnc[3];
	melhumeroj:=mnj[0]+mnj[1]+mnj[2]+mnj[3];
	if melhumeroc<melhumeroj{
		pontosj++;
		fmt.Print("Jogador obteve o melhor número. [");
		fmt.Print(melhumeroj);
		fmt.Print("-");
		fmt.Print(melhumeroc);
		fmt.Println("]");
	}else if melhumeroc>melhumeroj{
		pontosc++;
		fmt.Println("Computador obteve o melhor número. [");
		fmt.Print(melhumeroc);
		fmt.Print("-");
		fmt.Print(melhumeroj);
		fmt.Println("]");
	}/*else{
		fmt.Println("Ambos obtiveram o mesmo \"melhor número\".");
	}*/
	pontosj+=escopasjogador;
	pontosc+=escopascomputador;
	if pontosc<pontosj{
		fmt.Print("Jogador ganhou. [");
		fmt.Print(pontosj);
		fmt.Print("-");
		fmt.Print(pontosc);
		fmt.Println("]");
	}else if pontosc>pontosj{
		fmt.Print("Computador ganhou. [");
		fmt.Print(pontosc);
		fmt.Print("-");
		fmt.Print(pontosj);
		fmt.Println("]");
	}else{
		fmt.Print("Houve um empate. [");
		fmt.Print(pontosc);
		fmt.Print("-");
		fmt.Print(pontosj);
		fmt.Println("]");
	}
	time.Sleep(20*time.Second);
}