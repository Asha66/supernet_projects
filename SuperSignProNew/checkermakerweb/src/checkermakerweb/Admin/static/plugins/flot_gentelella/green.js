var theme = {
                    // é»˜è®¤è‰²æ¿
                    color: [
                        '#26B99A', '#34495E', '#BDC3C7', '#3498DB',
                        '#9B59B6', '#8abb6f', '#759c6a', '#bfd3b7'
                    ],
                    // å›¾è¡¨æ ‡é¢˜
                    title: {
                        itemGap: 8,
                        textStyle: {
                            fontWeight: 'normal',
                            color: '#408829'
                        }
                    },
                    // å€¼åŸŸ
                    dataRange: {
                        color: ['#1f610a', '#97b58d']
                    },
                    // å·¥å…·ç®±
                    toolbox: {
                        color: ['#408829', '#408829', '#408829', '#408829']
                    },
                    // æç¤ºæ¡†
                    tooltip: {
                        backgroundColor: 'rgba(0,0,0,0.5)',
                        axisPointer: {// åæ ‡è½´æŒ‡ç¤ºå™¨ï¼Œåæ ‡è½´è§¦å‘æœ‰æ•ˆ
                            type: 'line', // é»˜è®¤ä¸ºç›´çº¿ï¼Œå¯é€‰ä¸ºï¼š'line' | 'shadow'
                            lineStyle: {// ç›´çº¿æŒ‡ç¤ºå™¨æ ·å¼è®¾ç½®
                                color: '#408829',
                                type: 'dashed'
                            },
                            crossStyle: {
                                color: '#408829'
                            },
                            shadowStyle: {// é˜´å½±æŒ‡ç¤ºå™¨æ ·å¼è®¾ç½®
                                color: 'rgba(200,200,200,0.3)'
                            }
                        }
                    },
                    // åŒºåŸŸç¼©æ”¾æŽ§åˆ¶å™¨
                    dataZoom: {
                        dataBackgroundColor: '#eee', // æ•°æ®èƒŒæ™¯é¢œè‰²
                        fillerColor: 'rgba(64,136,41,0.2)', // å¡«å……é¢œè‰²
                        handleColor: '#408829'     // æ‰‹æŸ„é¢œè‰²
                    },
                    grid: {
                        borderWidth: 0
                    },
                    // ç±»ç›®è½´
                    categoryAxis: {
                        axisLine: {// åæ ‡è½´çº¿
                            lineStyle: {// å±žæ€§lineStyleæŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: '#408829'
                            }
                        },
                        splitLine: {// åˆ†éš”çº¿
                            lineStyle: {// å±žæ€§lineStyleï¼ˆè¯¦è§lineStyleï¼‰æŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: ['#eee']
                            }
                        }
                    },
                    // æ•°å€¼åž‹åæ ‡è½´é»˜è®¤å‚æ•°
                    valueAxis: {
                        axisLine: {// åæ ‡è½´çº¿
                            lineStyle: {// å±žæ€§lineStyleæŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: '#408829'
                            }
                        },
                        splitArea: {
                            show: true,
                            areaStyle: {
                                color: ['rgba(250,250,250,0.1)', 'rgba(200,200,200,0.1)']
                            }
                        },
                        splitLine: {// åˆ†éš”çº¿
                            lineStyle: {// å±žæ€§lineStyleï¼ˆè¯¦è§lineStyleï¼‰æŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: ['#eee']
                            }
                        }
                    },
                    timeline: {
                        lineStyle: {
                            color: '#408829'
                        },
                        controlStyle: {
                            normal: {color: '#408829'},
                            emphasis: {color: '#408829'}
                        }
                    },
                    // Kçº¿å›¾é»˜è®¤å‚æ•°
                    k: {
                        itemStyle: {
                            normal: {
                                color: '#68a54a', // é˜³çº¿å¡«å……é¢œè‰²
                                color0: '#a9cba2', // é˜´çº¿å¡«å……é¢œè‰²
                                lineStyle: {
                                    width: 1,
                                    color: '#408829', // é˜³çº¿è¾¹æ¡†é¢œè‰²
                                    color0: '#86b379'   // é˜´çº¿è¾¹æ¡†é¢œè‰²
                                }
                            }
                        }
                    },
                    map: {
                        itemStyle: {
                            normal: {
                                areaStyle: {
                                    color: '#ddd'
                                },
                                label: {
                                    textStyle: {
                                        color: '#c12e34'
                                    }
                                }
                            },
                            emphasis: {// ä¹Ÿæ˜¯é€‰ä¸­æ ·å¼
                                areaStyle: {
                                    color: '#99d2dd'
                                },
                                label: {
                                    textStyle: {
                                        color: '#c12e34'
                                    }
                                }
                            }
                        }
                    },
                    force: {
                        itemStyle: {
                            normal: {
                                linkStyle: {
                                    strokeColor: '#408829'
                                }
                            }
                        }
                    },
                    chord: {
                        padding: 4,
                        itemStyle: {
                            normal: {
                                lineStyle: {
                                    width: 1,
                                    color: 'rgba(128, 128, 128, 0.5)'
                                },
                                chordStyle: {
                                    lineStyle: {
                                        width: 1,
                                        color: 'rgba(128, 128, 128, 0.5)'
                                    }
                                }
                            },
                            emphasis: {
                                lineStyle: {
                                    width: 1,
                                    color: 'rgba(128, 128, 128, 0.5)'
                                },
                                chordStyle: {
                                    lineStyle: {
                                        width: 1,
                                        color: 'rgba(128, 128, 128, 0.5)'
                                    }
                                }
                            }
                        }
                    },
                    gauge: {
                        startAngle: 225,
                        endAngle: -45,
                        axisLine: {// åæ ‡è½´çº¿
                            show: true, // é»˜è®¤æ˜¾ç¤ºï¼Œå±žæ€§showæŽ§åˆ¶æ˜¾ç¤ºä¸Žå¦
                            lineStyle: {// å±žæ€§lineStyleæŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: [[0.2, '#86b379'], [0.8, '#68a54a'], [1, '#408829']],
                                width: 8
                            }
                        },
                        axisTick: {// åæ ‡è½´å°æ ‡è®°
                            splitNumber: 10, // æ¯ä»½splitç»†åˆ†å¤šå°‘æ®µ
                            length: 12, // å±žæ€§lengthæŽ§åˆ¶çº¿é•¿
                            lineStyle: {// å±žæ€§lineStyleæŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: 'auto'
                            }
                        },
                        axisLabel: {// åæ ‡è½´æ–‡æœ¬æ ‡ç­¾ï¼Œè¯¦è§axis.axisLabel
                            textStyle: {// å…¶ä½™å±žæ€§é»˜è®¤ä½¿ç”¨å…¨å±€æ–‡æœ¬æ ·å¼ï¼Œè¯¦è§TEXTSTYLE
                                color: 'auto'
                            }
                        },
                        splitLine: {// åˆ†éš”çº¿
                            length: 18, // å±žæ€§lengthæŽ§åˆ¶çº¿é•¿
                            lineStyle: {// å±žæ€§lineStyleï¼ˆè¯¦è§lineStyleï¼‰æŽ§åˆ¶çº¿æ¡æ ·å¼
                                color: 'auto'
                            }
                        },
                        pointer: {
                            length: '90%',
                            color: 'auto'
                        },
                        title: {
                            textStyle: {// å…¶ä½™å±žæ€§é»˜è®¤ä½¿ç”¨å…¨å±€æ–‡æœ¬æ ·å¼ï¼Œè¯¦è§TEXTSTYLE
                                color: '#333'
                            }
                        },
                        detail: {
                            textStyle: {// å…¶ä½™å±žæ€§é»˜è®¤ä½¿ç”¨å…¨å±€æ–‡æœ¬æ ·å¼ï¼Œè¯¦è§TEXTSTYLE
                                color: 'auto'
                            }
                        }
                    },
                    textStyle: {
                        fontFamily: 'å¾®è½¯é›…é»‘, Arial, Verdana, sans-serif'
                    }
                }